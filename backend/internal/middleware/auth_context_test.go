package middleware

import (
	"context"
	"testing"
)

func TestCancelledRequestStopsAuthLookup(t *testing.T) {
	var cache authContextCache
	first, cancel := context.WithCancel(context.Background())
	cancel()
	if cache.Next(first).Err() != context.Canceled {
		t.Fatal("first request cancellation was detached")
	}
	second := context.WithValue(context.Background(), authContextKey("request"), "second")
	resolved := cache.Next(second)
	if resolved.Err() != nil || resolved.Value(authContextKey("request")) != "second" {
		t.Fatalf("auth reused canceled request context: err=%v value=%v", resolved.Err(), resolved.Value(authContextKey("request")))
	}
}

type authContextKey string
