package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

const DefaultQuestion = "What kind of request is this?"
const MaxTiers = 10

type Tier struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Models      []string `yaml:"models"`
}

type Config struct {
	Classifier struct {
		Question    string  `yaml:"question"`
		Temperature float64 `yaml:"temperature"`
	} `yaml:"classifier"`
	DefaultTier string `yaml:"default_tier"`
	Tiers       []Tier `yaml:"tiers"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}
	if config.Classifier.Temperature == 0 {
		config.Classifier.Temperature = 0
	}
	return &config, config.validate()
}

func (config *Config) validate() error {
	if len(config.Tiers) == 0 {
		return errors.New("at least one tier is required")
	}
	if len(config.Tiers) > MaxTiers {
		return fmt.Errorf("tier count %d over max limit of %d", len(config.Tiers), MaxTiers)
	}
	seen := map[string]bool{}
	for i, tier := range config.Tiers {
		switch {
		case tier.Name == "":
			return fmt.Errorf("tiers[%d]: name is required", i)
		case seen[tier.Name]:
			return fmt.Errorf("tiers[%d]: duplicate name %q", i, tier.Name)
		case tier.Description == "" && len(config.Tiers) > 1:
			return fmt.Errorf("tier %q: description is required", tier.Name)
		case len(tier.Models) == 0:
			return fmt.Errorf("tier %q: at least one model is required", tier.Name)
		}
		seen[tier.Name] = true
	}
	if config.DefaultTier != "" && !seen[config.DefaultTier] {
		return fmt.Errorf("default_tier %q is not a defined tier", config.DefaultTier)
	}
	if config.Classifier.Question == "" {
		return fmt.Errorf("default_question is not defined")
	}
	if _, ok := seen[config.DefaultTier]; !ok {
		return fmt.Errorf("default tier is not defined")
	}
	return nil
}
