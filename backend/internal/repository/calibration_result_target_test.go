package repository

import (
	"errors"
	"testing"
)

func TestCalibrationMissingRecordKeepsErrorIdentity(t *testing.T) {
	db := newTestDB(t)
	repo := NewCalibrationRepository(db)
	_, err := repo.FindResultTarget(db, 404)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing record identity lost: %v", err)
	}
}
