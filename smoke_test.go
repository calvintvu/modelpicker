package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"runtime"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/calvinvu/ez-model-router/config"
	"github.com/calvinvu/ez-model-router/inference"
)

func TestRealModelHTTP(t *testing.T) {
	baseURL := strings.TrimRight(os.Getenv("ROUTER_SMOKE_URL"), "/")
	if baseURL == "" {
		t.Skip("set ROUTER_SMOKE_URL to run the explicit real-model HTTP check")
	}
	key := os.Getenv("API_KEY")
	if key == "" {
		t.Fatal("API_KEY is required for the smoke check")
	}
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "testdata/smoke.yaml"
	}
	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Tiers) < 2 || cfg.DefaultTier != "" {
		t.Fatal("smoke configuration needs at least two tiers and no default_tier, so inference errors cannot pass as fallback")
	}
	healthClient := &http.Client{Timeout: 5 * time.Second}
	defer healthClient.CloseIdleConnections()
	deadline := time.Now().Add(180 * time.Second)
	for {
		response, err := healthClient.Get(baseURL + "/healthz")
		if err == nil {
			body, readErr := io.ReadAll(io.LimitReader(response.Body, 1024))
			response.Body.Close()
			if response.StatusCode == http.StatusOK && readErr == nil && len(body) == 0 {
				break
			}
		}
		if time.Now().After(deadline) {
			t.Fatal("health endpoint did not return HTTP 200 with an empty body before the startup deadline")
		}
		time.Sleep(200 * time.Millisecond)
	}
	client := &http.Client{Timeout: 120 * time.Second}
	defer client.CloseIdleConnections()
	for _, authenticated := range []bool{false, true} {
		request, err := http.NewRequest(http.MethodPost, baseURL+"/v1/route", strings.NewReader(`{"prompt":"Explain how a Go function works."}`))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		if authenticated {
			request.Header.Set("Authorization", "Bearer "+key)
		}
		response, err := client.Do(request)
		if err != nil {
			t.Fatal(err)
		}
		if !authenticated {
			response.Body.Close()
			if response.StatusCode != http.StatusUnauthorized || response.Header.Get("WWW-Authenticate") != "Bearer" {
				t.Fatal("unauthenticated route did not return HTTP 401 with a bearer challenge")
			}
			continue
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK || response.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("authenticated route returned HTTP %d or an unexpected content type", response.StatusCode)
		}
		var result inference.ModelRouteResponse
		decoder := json.NewDecoder(io.LimitReader(response.Body, maxBodyBytes))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&result); err != nil {
			t.Fatal(err)
		}
		if err := decoder.Decode(&struct{}{}); err != io.EOF {
			t.Fatal("route response contains trailing JSON")
		}
		index := slices.IndexFunc(cfg.Tiers, func(tier config.Tier) bool { return tier.Name == result.Tier })
		if index < 0 || !slices.Contains(cfg.Tiers[index].Models, result.Model) {
			t.Fatal("route response model is not a member of its configured tier")
		}
	}
	t.Logf("real-model health/auth/route membership passed; client target %s/%s", runtime.GOOS, runtime.GOARCH)
}
