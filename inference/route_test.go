package inference

import (
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/calvinvu/ez-model-router/config"
)

func TestRouteOneTier(t *testing.T) {
	cfg := &config.Config{Tiers: []config.Tier{{Name: "only", Models: []string{"a", "b"}}}}
	for range 20 {
		result, err := Route(httptest.NewRequest("POST", "/v1/route", nil), cfg, "hello", nil)
		if err != nil || result.Tier != "only" || !slices.Contains(cfg.Tiers[0].Models, result.Model) {
			t.Fatalf("invalid one-tier result: %+v, %v", result, err)
		}
	}
}
