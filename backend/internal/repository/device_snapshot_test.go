package repository

import (
	"testing"
	"time"

	"github.com/medasset/medasset/internal/model"
)

func TestDeviceRefreshKeepsStableOrdering(t *testing.T) {
	expiry := time.Date(2028, 6, 1, 0, 0, 0, 0, time.UTC)
	input := []model.Device{{ID: 1, WarrantyExpiry: &expiry}, {ID: 3}, {ID: 2}}
	snapshot := SnapshotDeviceList(input)
	if input[0].ID != 1 || snapshot[0].ID != 3 {
		t.Fatalf("snapshot ordering changed input: input=%v snapshot=%v", ids(input), ids(snapshot))
	}
	*snapshot[2].WarrantyExpiry = snapshot[2].WarrantyExpiry.AddDate(1, 0, 0)
	if expiry.Year() != 2028 {
		t.Fatalf("snapshot retained nested pointer: %v", expiry)
	}
}

func ids(devices []model.Device) []uint {
	result := make([]uint, len(devices))
	for i := range devices {
		result[i] = devices[i].ID
	}
	return result
}
