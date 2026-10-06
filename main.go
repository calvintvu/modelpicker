package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/calvinvu/ez-model-router/inference"
)

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
		routeResponse := Route()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routeResponse)
	})

	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
