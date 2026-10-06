package main

type ModelRouteResponse struct {
	Model string `json:"model"`
}

func Route() ModelRouteResponse {
	return ModelRouteResponse{Model: "fake-model-slug"}
}
