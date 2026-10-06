package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/calvinvu/ez-model-router/inference"
)

const maxBodyBytes = 1 << 20

func main() {
	session, err := inference.LoadInferenceSession()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load inference session:", err)
		os.Exit(1)
	}
	defer session.CloseInferenceSession()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("POST /v1/route", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
		var req ModelRouteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		if req.Prompt == "" {
			http.Error(w, "prompt is required", http.StatusBadRequest)
			return
		}
		routeResponse, err := Route(req.Prompt, session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routeResponse)
	})

	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
