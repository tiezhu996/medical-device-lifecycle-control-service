package pointerx_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/medasset/medasset/internal/config"
	"github.com/medasset/medasset/internal/router"
)

func TestPreflightResponseCarriesRequestID(t *testing.T) {
	cfg := &config.Config{RunMode: "test", JWTSecret: "test-secret", RateLimit: 100}
	engine := router.New(router.Deps{Cfg: cfg, Log: slog.Default()})
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodOptions, "/api/v1/auth/login", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("preflight status = %d", w.Code)
	}
	if got := w.Header().Get("X-Request-Id"); got == "" {
		t.Fatal("preflight response omitted request id")
	}
}
