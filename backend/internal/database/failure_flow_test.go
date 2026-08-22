package database_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/medasset/medasset/internal/util"
)

func TestFailureResponseStopsHandlerChain(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	continued := false
	r.GET("/protected",
		func(c *gin.Context) {
			util.Fail(c, http.StatusForbidden, 40300, "access denied")
		},
		func(c *gin.Context) {
			continued = true
			util.OK(c, gin.H{"unexpected": true})
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
