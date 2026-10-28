package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/calvinvu/ez-model-router/config"
)

func TestHTTPContract(t *testing.T) {
	cfg := &config.Config{Tiers: []config.Tier{{Name: "only", Models: []string{"test/only"}}}}
	server := httptest.NewServer(newHandler("test-key", cfg, nil))
	defer server.Close()
	for _, tc := range []struct {
		name, method, path, body, authorization, wantBody string
		status                                            int
	}{
		{name: "health without auth", method: "GET", path: "/healthz", status: 200},
		{name: "head health", method: "HEAD", path: "/healthz", status: 200},
		{name: "missing auth", method: "POST", path: "/v1/route", body: `{"prompt":"hi"}`, status: 401, wantBody: "unauthorized request\n"},
		{name: "empty body", body: "", status: 400, wantBody: "invalid JSON body\n"},
		{name: "malformed JSON", body: `{`, status: 400, wantBody: "invalid JSON body\n"},
		{name: "wrong prompt type", body: `{"prompt":23}`, status: 400, wantBody: "invalid JSON body\n"},
		{name: "missing prompt", body: `{}`, status: 400, wantBody: "prompt is required\n"},
		{name: "null prompt", body: `{"prompt":null}`, status: 400, wantBody: "prompt is required\n"},
		{name: "empty prompt", body: `{"prompt":""}`, status: 400, wantBody: "prompt is required\n"},
		{name: "valid", body: `{"prompt":"hello"}`, status: 200},
		{name: "whitespace prompt", body: `{"prompt":" "}`, status: 200},
		{name: "unknown fields ignored", body: `{"prompt":"hi","extra":true}`, status: 200},
		{name: "trailing object", body: `{"prompt":"hi"}{}`, status: 400, wantBody: "invalid JSON body\n"},
		{name: "trailing garbage", body: `{"prompt":"hi"}garbage`, status: 400, wantBody: "invalid JSON body\n"},
		{name: "oversized prompt", body: `{"prompt":"` + strings.Repeat("x", maxBodyBytes) + `"}`, status: 400, wantBody: "invalid JSON body\n"},
		{name: "oversized trailing whitespace", body: `{"prompt":"hi"}` + strings.Repeat(" ", maxBodyBytes), status: 400, wantBody: "invalid JSON body\n"},
		{name: "exact body limit", body: `{"prompt":"` + strings.Repeat("x", maxBodyBytes-len(`{"prompt":""}`)) + `"}`, status: 200},
		{name: "wrong method", method: "GET", path: "/v1/route", status: 405, wantBody: "Method Not Allowed\n"},
		{name: "unknown path", method: "GET", path: "/missing", status: 404, wantBody: "404 page not found\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			method, path, auth := tc.method, tc.path, tc.authorization
			if method == "" {
				method = http.MethodPost
				path = "/v1/route"
				auth = "Bearer test-key"
			}
			request, err := http.NewRequest(method, server.URL+path, strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			if auth != "" {
				request.Header.Set("Authorization", auth)
			}
			response, err := server.Client().Do(request)
			if err != nil {
				t.Fatal(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}
			if response.StatusCode != tc.status {
				t.Fatalf("HTTP %d, want %d: %s", response.StatusCode, tc.status, body)
			}
			if tc.status == 200 && path == "/v1/route" {
				var result map[string]string
				if err := json.Unmarshal(body, &result); err != nil {
					t.Fatal(err)
				}
				if len(result) != 2 || result["tier"] != "only" || result["model"] != "test/only" || response.Header.Get("Content-Type") != "application/json" {
					t.Fatalf("unexpected route response: %s", body)
				}
			} else if string(body) != tc.wantBody {
				t.Fatalf("body %q, want %q", body, tc.wantBody)
			}
			if tc.status >= 400 && response.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
				t.Fatal("error response is not plain text")
			}
			if tc.status == 401 && response.Header.Get("WWW-Authenticate") != "Bearer" {
				t.Fatal("missing bearer challenge")
			}
		})
	}
}
