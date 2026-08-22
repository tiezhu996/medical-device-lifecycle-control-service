package middleware

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRateLimitWaitHonorsRequestDeadline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := time.Now()
	err := waitForRateLimit(ctx, 150*time.Millisecond)
	if !errors.Is(err, context.Canceled) || time.Since(started) > 100*time.Millisecond {
		t.Fatalf("rate limit wait ignored cancellation: err=%v elapsed=%v", err, time.Since(started))
	}
}
