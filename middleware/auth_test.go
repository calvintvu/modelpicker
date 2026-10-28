package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequireAPIKey(t *testing.T) {
	var logs bytes.Buffer
	oldOutput := log.Writer()
	log.SetOutput(&logs)
	t.Cleanup(func() { log.SetOutput(oldOutput) })
	for _, tc := range []struct {
		header string
		ok     bool
	}{
		{"", false}, {"Basic private-candidate", false}, {"Bearer wrong-private-candidate", false},
		{"Bearer", false}, {"Bearer  private-expected-key", false},
		{"Bearer private-expected-key", true}, {"bearer private-expected-key", true},
	} {
		called := false
		handler := RequireAPIKey("private-expected-key")(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			called = true
			w.WriteHeader(http.StatusNoContent)
		}))
		r := httptest.NewRequest("POST", "/v1/pick", nil)
		r.Header.Set("Authorization", tc.header)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		if called != tc.ok {
			t.Fatalf("authorization decision: got %v, want %v", called, tc.ok)
		}
		if !tc.ok && (w.Code != 401 || w.Header().Get("WWW-Authenticate") != "Bearer" || w.Body.String() != "unauthorized request\n") {
			t.Fatal("incorrect unauthorized response")
		}
	}
	if strings.Contains(logs.String(), "private-") {
		t.Fatal("authentication logs exposed an expected or supplied key")
	}
}
