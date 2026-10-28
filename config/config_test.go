package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfigRejectsUnusableSelectionValues(t *testing.T) {
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
				t.Fatal("expected invalid selection configuration to be rejected")
			}
		})
	}
}

func TestSelectionConfigurationBoundaries(t *testing.T) {
	for _, tc := range []struct{ name, yaml, want string }{
		{"unknown field", "unknown: true\ntiers: [{name: a, models: [m]}]", "field unknown"},
		{"no tiers", "tiers: []", "at least one tier"},
		{"duplicate names", "tiers: [{name: a, description: first, models: [m]}, {name: a, description: second, models: [m]}]", "duplicate name"},
		{"no models", "tiers: [{name: a, models: []}]", "at least one model"},
		{"unknown fallback", "default_tier: missing\ntiers: [{name: a, models: [m]}]", "not a defined tier"},
		{"negative temperature", "classifier: {temperature: -1}\ntiers: [{name: a, models: [m]}]", "temperature must be finite"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.yaml")
			if err := os.WriteFile(path, []byte(tc.yaml), 0600); err != nil {
				t.Fatal(err)
			}
			if _, err := LoadConfig(path); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("expected %q, got %v", tc.want, err)
			}
		})
	}
	for _, count := range []int{MaxTiers, MaxTiers + 1} {
		cfg := &Config{}
		for i := range count {
			cfg.Tiers = append(cfg.Tiers, Tier{Name: string(rune('a' + i)), Description: "kind", Models: []string{"m"}})
		}
		if err := cfg.validate(); (err == nil) != (count == MaxTiers) {
			t.Fatalf("tier count %d: %v", count, err)
		}
	}
}

func TestValidateFit(t *testing.T) {
	cfg := &Config{Tiers: []Tier{{Name: "only", Models: []string{"m"}}}}
	if err := cfg.ValidateFit(func(string, []string) error {
		t.Fatal("one-tier configuration must bypass classifier fit checks")
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	cfg.Tiers = append(cfg.Tiers, Tier{Name: "other", Description: "other kind", Models: []string{"n"}})
	fitError := errors.New("token budget exceeded")
	if err := cfg.ValidateFit(func(string, []string) error { return fitError }); !errors.Is(err, fitError) {
		t.Fatalf("configuration must preserve the fit error, got %v", err)
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
