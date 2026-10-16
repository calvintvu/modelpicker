package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/calvinvu/ez-model-router/config"
	"github.com/calvinvu/ez-model-router/inference"
	"github.com/calvinvu/ez-model-router/middleware"
	"github.com/joho/godotenv"
	"go.yaml.in/yaml/v3"
)

const maxBodyBytes = 1 << 20

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to load .env:", err)
		os.Exit(1)
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "API_KEY is not set")
		os.Exit(1)
	}

	session, err := inference.LoadInferenceSession()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load inference session:", err)
		os.Exit(1)
	}
	defer session.CloseInferenceSession()

	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid config:", err)
		os.Exit(1)
	}
	if err := cfg.ValidateFit(session.CheckFits); err != nil {
		fmt.Fprintln(os.Stderr, "invalid config:", err)
		os.Exit(1)
	}
	out, _ := yaml.Marshal(cfg)
	fmt.Println(string(out))

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.Handle("POST /v1/route", middleware.RequireAPIKey(apiKey)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
