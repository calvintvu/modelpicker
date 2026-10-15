package inference

import (
	"log"

	"github.com/calvinvu/ez-model-router/config"
)

type ModelRouteRequest struct {
	Prompt string `json:"prompt"`
}

type ModelRouteResponse struct {
	Tier string `json:"tier"`
}

// calibrated from laya_config.json
func defaultTemperature(choiceCount int) float64 {
	switch {
	case choiceCount <= 2:
		return 1.9063563346862793
	case choiceCount <= 5:
		return 1.7601518630981445
	case choiceCount <= 10:
		return 1.0000158548355103
	default:
		return 0.10058280825614929
	}
}

func Route(cfg *config.Config, prompt string, session *InferenceSession) (ModelRouteResponse, error) {
	options := make([]string, len(cfg.Tiers))
	for i, tier := range cfg.Tiers {
		options[i] = tier.Name + ": " + tier.Description
	}
	temperature := cfg.Classifier.Temperature
	if temperature == 0 {
		temperature = defaultTemperature(len(options))
	}

	probs, err := session.Choose(cfg.Classifier.Question, prompt, options, temperature)
	if err != nil {
		return ModelRouteResponse{}, err
	}
	best := 0
	for i := range probs {
		if probs[i] > probs[best] {
			best = i
		}
	}
	tier := cfg.Tiers[best].Name
	log.Println(tier)
	return ModelRouteResponse{Tier: tier}, nil
}
