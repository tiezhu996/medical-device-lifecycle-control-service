package util

import (
	"context"
	"errors"
	"testing"
)

func TestNextRequestDoesNotReuseCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := NewContextError(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation identity was lost in middleware error: %v", err)
	}
}
