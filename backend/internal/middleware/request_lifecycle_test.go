package middleware

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/medasset/medasset/internal/util"
)

func TestRequestIDVisibleBeforeHandlerRuns(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RequestID())
	var seen any
	r.GET("/devices", func(c *gin.Context) {
		seen, _ = c.Get(RequestIDKey)
		c.Status(http.StatusNoContent)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	req.Header.Set("X-Request-Id", "device-read-42")
	r.ServeHTTP(w, req)

	if seen != "device-read-42" {
		t.Fatalf("handler saw request id %q", seen)
	}
	if got := w.Header().Get("X-Request-Id"); got != "device-read-42" {
		t.Fatalf("response request id = %q", got)
	}
}

func TestPanicRecoveryProducesSingleErrorPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RequestID())
	r.Use(ErrorHandler(slog.New(slog.NewTextHandler(io.Discard, nil))))
	r.GET("/devices", func(c *gin.Context) {
		panic("device adapter unavailable")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	r.ServeHTTP(w, req)

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
