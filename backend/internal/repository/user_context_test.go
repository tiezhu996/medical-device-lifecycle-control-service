package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/medasset/medasset/internal/model"
)

func TestLoginCancellationReachesRepository(t *testing.T) {
	repo := NewUserRepository(newTestDB(t))
	if err := repo.Create(&model.User{Username: "ctx-user", Password: "hash", RealName: "Context", Role: "DEPARTMENT", Status: "active"}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := repo.FindByUsernameContext(ctx, "ctx-user")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("repository query ignored cancellation: %v", err)
	}
}
