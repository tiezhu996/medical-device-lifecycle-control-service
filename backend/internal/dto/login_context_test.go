package dto

import (
	"context"
	"testing"
)

func TestRegisterUsesCallerContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if ResolveLoginContext(ctx).Err() != context.Canceled {
		t.Fatal("request cancellation was detached while resolving login context")
	}
}
