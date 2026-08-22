package dto

import (
	"testing"
	"time"

	"github.com/medasset/medasset/internal/model"
)

func TestDeviceSnapshotsIsolateConcurrentRefresh(t *testing.T) {
	expiry := time.Date(2027, 1, 2, 0, 0, 0, 0, time.UTC)
	device := &model.Device{ID: 7, WarrantyExpiry: &expiry}
	detail := NewDeviceDetail(device, false)
	*device.WarrantyExpiry = device.WarrantyExpiry.AddDate(1, 0, 0)
	if detail.WarrantyExpiry.Year() != 2027 {
		t.Fatalf("detail snapshot changed with source: %v", detail.WarrantyExpiry)
	}
}
