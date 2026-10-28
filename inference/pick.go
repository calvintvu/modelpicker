package inference

import (
	"log"
	"net/http"
	"slices"

	"github.com/calvintvu/modelpicker/config"
	random "github.com/calvintvu/modelpicker/utils"
)

type ModelPickRequest struct {
	Prompt string `json:"prompt"`
}

type ModelPickResponse struct {
	Tier  string `json:"tier"`
	Model string `json:"model"`
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

func Pick(r *http.Request, cfg *config.Config, prompt string, session *InferenceSession) (ModelPickResponse, error) {
	if len(cfg.Tiers) == 1 {
		return pick(cfg.Tiers[0])
	}
	options := cfg.Options()
	temperature := cfg.Classifier.Temperature
	if temperature == 0 {
		temperature = defaultTemperature(len(options))
	}

	probs, err := session.Choose(cfg.Classifier.Question, prompt, options, temperature)
	if err != nil {
		i := slices.IndexFunc(cfg.Tiers, func(t config.Tier) bool { return t.Name == cfg.DefaultTier })
		if i < 0 {
			return ModelPickResponse{}, err
		}
		log.Printf("inference failed, using default tier %q: %v", cfg.DefaultTier, err)
		return pick(cfg.Tiers[i])
	}
	best := 0
	for i := range probs {
		if probs[i] > probs[best] {
			best = i
		}
	}
	log.Println(r.URL.Path, cfg.Tiers[best].Name)
	return pick(cfg.Tiers[best])
}

func pick(tier config.Tier) (ModelPickResponse, error) {
	model, err := random.RandomElement(tier.Models)
	if err != nil {
		return ModelPickResponse{}, err
	}
	return ModelPickResponse{Tier: tier.Name, Model: model}, nil
}
