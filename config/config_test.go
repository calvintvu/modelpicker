package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigRejectsUnusableRoutingValues(t *testing.T) {
	for _, tc := range []struct{ name, yaml string }{
		{"blank model", "tiers: [{name: only, models: [' ']}]"},
		{"blank tier name", "tiers: [{name: ' ', models: [model]}]"},
		{"blank description", "tiers: [{name: a, description: ' ', models: [a]}, {name: b, description: other, models: [b]}]"},
		{"NaN temperature", "classifier: {temperature: .nan}\ntiers: [{name: only, models: [model]}]"},
		{"infinite temperature", "classifier: {temperature: .inf}\ntiers: [{name: only, models: [model]}]"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.yaml")
			if err := os.WriteFile(path, []byte(tc.yaml), 0600); err != nil {
				t.Fatal(err)
			}
			if _, err := LoadConfig(path); err == nil {
				t.Fatal("expected invalid routing configuration to be rejected")
			}
		})
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte("tiers: [{name: only, models: [model]}]"), 0600); err != nil {
		t.Fatal(err)
	}
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Classifier.Question != DefaultQuestion || cfg.Classifier.Temperature != 0 {
		t.Fatalf("unexpected classifier defaults: %+v", cfg.Classifier)
	}
}
