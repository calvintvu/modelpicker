package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/calvinvu/ez-model-router/config"
	"github.com/calvinvu/ez-model-router/inference"
	"github.com/calvinvu/ez-model-router/middleware"
	"github.com/joho/godotenv"
	"go.yaml.in/yaml/v3"
)

const (
	maxBodyBytes   = 1 << 20
	maxHeaderBytes = 32 << 10
)

func main() {
	if err := runServer(); err != nil {
		log.Printf("fatal: %v", err)
		os.Exit(1)
	}
}

func runServer() error {
	if _, err := os.Lstat(".env"); !errors.Is(err, os.ErrNotExist) {
		if err := godotenv.Load(); err != nil {
			var pathErr *os.PathError
			if errors.As(err, &pathErr) {
				return fmt.Errorf("failed to load .env: %w", err)
			}
			// godotenv parse errors can contain secrets
			return errors.New("failed to load .env: invalid syntax")
		}
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return errors.New("API_KEY is not set")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yaml"
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("invalid config %q: %w", configPath, err)
	}

	session, err := inference.LoadInferenceSession()
	if err != nil {
		return fmt.Errorf("failed to load inference session: %w", err)
	}
	defer session.CloseInferenceSession()
	if err := cfg.ValidateFit(session.CheckFits); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}
	out, _ := yaml.Marshal(cfg)
	log.Printf("loaded config:\n%s", out)

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:              listenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("POST /v1/route", middleware.RequireAPIKey(apiKey)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
		var req inference.ModelRouteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		if req.Prompt == "" {
			http.Error(w, "prompt is required", http.StatusBadRequest)
			return
		}
		routeResponse, err := inference.Route(r, cfg, req.Prompt, session)
		if err != nil {
			log.Printf("route failed: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routeResponse)
	})))

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("server starting on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	var exitErr error
	select {
	case err := <-serverErrors:
		exitErr = fmt.Errorf("server error: %w", err)
	case sig := <-quit:
		log.Printf("received signal: %v", sig)
	}

	log.Println("server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
		if closeErr := server.Close(); closeErr != nil {
			log.Printf("error closing server: %v", closeErr)
		}
	}

	log.Println("server stopped")
	return exitErr
}
