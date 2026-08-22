package model

import (
	"testing"
	"time"
)

func TestDeviceListReturnsDeepCopy(t *testing.T) {
	purchased := time.Date(2026, 2, 3, 0, 0, 0, 0, time.UTC)
	device := Device{ID: 2, PurchaseDate: &purchased}
	snapshot := CloneDevice(&device)
	*snapshot.PurchaseDate = snapshot.PurchaseDate.AddDate(0, 1, 0)
	if device.PurchaseDate.Month() != time.February {
		t.Fatalf("source date changed through snapshot: %v", device.PurchaseDate)
	}
}
