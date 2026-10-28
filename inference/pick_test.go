package inference

import (
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/calvintvu/modelpicker/config"
)

func TestPickOneTier(t *testing.T) {
	cfg := &config.Config{Tiers: []config.Tier{{Name: "only", Models: []string{"a", "b"}}}}
	for range 20 {
		result, err := Pick(httptest.NewRequest("POST", "/v1/pick", nil), cfg, "hello", nil)
		if err != nil || result.Tier != "only" || !slices.Contains(cfg.Tiers[0].Models, result.Model) {
			t.Fatalf("invalid one-tier result: %+v, %v", result, err)
		}
	}
}
