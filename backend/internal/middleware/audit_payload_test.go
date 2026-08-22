package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestAuditMiddlewarePublishesCompletePayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var detail string
	router := gin.New()
	router.PATCH("/devices/:id", func(c *gin.Context) {
		payload := buildAuditPayload(c, 7, "engineer", "req-77", time.Now())
		if payload.RequestID != "req-77" || payload.Module != "/devices/:id" || payload.EntityID != "31" {
			t.Fatalf("incomplete audit identity: %#v", payload)
		}
		detail = payload.Detail
	})
	req := httptest.NewRequest(http.MethodPatch, "/devices/31", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)
	if !strings.Contains(detail, "/devices/31") {
		t.Fatalf("audit detail omitted request path: %q", detail)
	}
}
