package util

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCancelledMiddlewareLeavesNoRetry(t *testing.T) {
	token, err := GenerateToken("ctx-secret", time.Hour, 3, "doctor", "DEPARTMENT")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = ParseTokenContext(ctx, "ctx-secret", token)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("token parsing continued after cancellation: %v", err)
	}
}
