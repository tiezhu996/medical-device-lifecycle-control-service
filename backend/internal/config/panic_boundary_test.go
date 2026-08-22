package config_test

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/medasset/medasset/internal/middleware"
	"github.com/medasset/medasset/internal/util"
)

func TestPanicRecoveryProducesSingleErrorPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.ErrorHandler(slog.New(slog.NewTextHandler(io.Discard, nil))))
	r.GET("/devices", func(c *gin.Context) {
		panic("device adapter unavailable")
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/devices", nil))

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("panic status = %d", w.Code)
	}
	var body util.Resp
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("panic response is not one JSON object: %v, body=%q", err, w.Body.String())
	}
	if body.Code == 0 {
		t.Fatalf("panic response used success code: %#v", body)
	}
}
