package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/medasset/medasset/internal/middleware"
)

func TestRequestIDVisibleBeforeHandlerRuns(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.RequestID())
	var seen any
	r.GET("/devices", func(c *gin.Context) {
		seen, _ = c.Get(middleware.RequestIDKey)
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
