package model

import (
	"context"
	"testing"
)

func TestLoginAuditStopsAfterCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	user := &User{ID: 8, Username: "operator"}
	if user.AuditContext(ctx).Err() != context.Canceled {
		t.Fatal("audit context ignored the canceled login request")
	}
}
