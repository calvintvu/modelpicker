package inference

type ModelRouteRequest struct {
	Prompt string `json:"prompt"`
}

type ModelRouteResponse struct {
	Tier string `json:"tier"`
}

func Route(prompt string, session *InferenceSession) (ModelRouteResponse, error) {
	questionText := "How difficult is this request to answer well?"
	options := []string{
		"complex: multi-step reasoning, code, math, analysis, long-form writing",
		"moderate: explanations, summaries, short writing tasks",
		"simple: greetings, small talk, short factual lookups",
	}
	tiers := []string{"Tier 1", "Tier 2", "Tier 3"}
	temperature := 1.7601518630981445 // choice:3-5 from laya_config.json

	probs, err := session.Choose(questionText, prompt, options, temperature)
	if err != nil {
		return ModelRouteResponse{}, err
	}
	best := 0
	for i := range probs {
		if probs[i] > probs[best] {
			best = i
		}
	}
	return ModelRouteResponse{Tier: tiers[best]}, nil
}
