package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestStartupConfiguration(t *testing.T) {
	for _, tc := range []struct {
		name, dotenv, apiKey, want string
		file, directory, external  bool
	}{
		{name: "missing dotenv", apiKey: "external-test-key", want: "invalid config"},
		{name: "missing API key", want: "API_KEY is not set"},
		{name: "dotenv API key", file: true, dotenv: "API_KEY=file-test-key\n", want: "invalid config"},
		{name: "external API key takes precedence", file: true, dotenv: "API_KEY=\n", apiKey: "external-test-key", want: "invalid config"},
		{name: "empty external API key", file: true, external: true, dotenv: "API_KEY=file-test-key\n", want: "API_KEY is not set"},
		{name: "malformed dotenv", file: true, dotenv: "API_KEY=\"unterminated-test-secret", apiKey: "external-test-key", want: "failed to load .env"},
		{name: "unreadable dotenv path", directory: true, apiKey: "external-test-key", want: "failed to load .env"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, key := range []string{"API_KEY", "CONFIG_PATH", "MODEL_DIR", "ONNX_RUNTIME_PATH", "LISTEN_ADDR", "WRITE_TIMEOUT"} {
				t.Setenv(key, "")
				if err := os.Unsetenv(key); err != nil {
					t.Fatal(err)
				}
			}
			dir := t.TempDir()
			t.Chdir(dir)
			if tc.file {
				if err := os.WriteFile(".env", []byte(tc.dotenv), 0600); err != nil {
					t.Fatal(err)
				}
			}
			if tc.directory {
				if err := os.Mkdir(".env", 0700); err != nil {
					t.Fatal(err)
				}
			}
			if tc.apiKey != "" || tc.external {
				t.Setenv("API_KEY", tc.apiKey)
			}
			configPath := filepath.Join(dir, "custom config.yaml")
			t.Setenv("CONFIG_PATH", configPath)
			t.Setenv("ONNX_RUNTIME_PATH", filepath.Join(dir, "missing-runtime"))
			err := runServer()
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("expected %q, got %v", tc.want, err)
			}
			if strings.Contains(err.Error(), "unterminated-test-secret") {
				t.Fatal("dotenv parse error exposed a secret value")
			}
			if tc.want == "invalid config" && !strings.Contains(err.Error(), configPath) {
				t.Fatalf("expected custom config path %q in error, got %v", configPath, err)
			}
			if tc.apiKey != "" && os.Getenv("API_KEY") != tc.apiKey {
				t.Fatal("dotenv replaced the external API key")
			}
		})
	}
}

func TestServerWriteTimeout(t *testing.T) {
	for _, tc := range []struct {
		value string
		want  time.Duration
	}{
		{"", time.Minute},
		{"90s", 90 * time.Second},
		{"250ms", 250 * time.Millisecond},
	} {
		t.Run(tc.value, func(t *testing.T) {
			t.Setenv("WRITE_TIMEOUT", tc.value)
			got, err := serverWriteTimeout()
			if err != nil || got != tc.want {
				t.Fatalf("got %v, %v; want %v", got, err, tc.want)
			}
		})
	}
	for _, value := range []string{"0", "-1s", "invalid-secret-value", "999999999999999999999h"} {
		t.Run(value, func(t *testing.T) {
			t.Setenv("WRITE_TIMEOUT", value)
			_, err := serverWriteTimeout()
			if err == nil || !strings.Contains(err.Error(), "WRITE_TIMEOUT") || strings.Contains(err.Error(), "invalid-secret-value") {
				t.Fatalf("expected actionable error without echoing the value, got %v", err)
			}
		})
	}
}

func TestStartupValidatesYAMLBeforeModelInitialization(t *testing.T) {
	t.Chdir(t.TempDir())
	t.Setenv("API_KEY", "test-key")
	t.Setenv("CONFIG_PATH", "")
	t.Setenv("ONNX_RUNTIME_PATH", "/missing-runtime")
	if err := os.WriteFile(".env", nil, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("config.yaml", []byte("unknown_field: true\n"), 0600); err != nil {
		t.Fatal(err)
	}
	err := runServer()
	if err == nil || !strings.Contains(err.Error(), "invalid config") || !strings.Contains(err.Error(), "unknown_field") {
		t.Fatalf("expected strict YAML validation before native initialization, got %v", err)
	}
}
