package repository

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

func TestPlanBatchPreservesWorkFailure(t *testing.T) {
	db := newTestDB(t)
	repo := NewMaintenanceRepository(db)
	want := errors.New("device schedule rejected")
	err := repo.RunPlanBatch(func(tx *gorm.DB) error {
		return want
	})
	if !errors.Is(err, want) {
		t.Fatalf("work failure lost: %v", err)
	}
}
