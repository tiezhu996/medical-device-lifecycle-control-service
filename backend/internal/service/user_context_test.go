package service

import (
	"context"
	"testing"
)

func TestLoginDeadlineDoesNotLeakAcrossRequests(t *testing.T) {
	service := &UserService{}
	first, cancel := context.WithCancel(context.Background())
	cancel()
	if service.nextAuthContext(first).Err() != context.Canceled {
		t.Fatal("first login did not retain its cancellation")
	}
	second := context.WithValue(context.Background(), contextKey("request"), "second")
	resolved := service.nextAuthContext(second)
	if resolved.Err() != nil || resolved.Value(contextKey("request")) != "second" {
		t.Fatalf("second login reused first request context: err=%v value=%v", resolved.Err(), resolved.Value(contextKey("request")))
	}
}

type contextKey string
