package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

const maxBodyBytes = 1 << 20

func main() {
	if err := runServer(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func runServer() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env: %w", err)
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return errors.New("API_KEY is not set")
	}

	session, err := inference.LoadInferenceSession()
	if err != nil {
		return fmt.Errorf("failed to load inference session: %w", err)
	}
	defer session.CloseInferenceSession()

	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}
	if err := cfg.ValidateFit(session.CheckFits); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}
	out, _ := yaml.Marshal(cfg)
	fmt.Println(string(out))

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routeResponse)
	})))

	serverErrors := make(chan error, 1)

	go func() {
		fmt.Println("Server starting on :8080")
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
		fmt.Println("Received signal:", sig)
	}

	fmt.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Graceful shutdown failed: %s\n", err)
		if closeErr := server.Close(); closeErr != nil {
			fmt.Printf("Error closing server: %s\n", closeErr)
		}
	}

	fmt.Println("Server stopped")
	return exitErr
}
