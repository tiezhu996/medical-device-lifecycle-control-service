package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFailureResponseStopsHandlerChain(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	continued := false
	r.GET("/protected",
		func(c *gin.Context) {
			Fail(c, http.StatusForbidden, 40300, "access denied")
		},
		func(c *gin.Context) {
			continued = true
			OK(c, gin.H{"unexpected": true})
		},
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/protected", nil))

	if continued {
		t.Fatal("business handler ran after failure response")
	}
	if w.Code != http.StatusForbidden {
		t.Fatalf("failure status = %d", w.Code)
	}
}
