package repository

import (
	"testing"
	"time"

	"github.com/medasset/medasset/internal/model"
)

func TestConditionalAcceptanceRejectsStaleStatus(t *testing.T) {
	db := newTestDB(t)
	p := &model.PurchaseRequest{
		RequestNo: "PR-STale", Department: "ICU", ApplicantName: "Li",
		DeviceName: "Monitor", Quantity: 1, Status: "approved",
	}
	if err := db.Create(p).Error; err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	p.ApplyAcceptance("engineer", &now, "probe", "CERT-1", "REG-1", 77)
	repo := NewPurchaseRepository(db)
	if err := repo.MarkAcceptedTx(db, p); err == nil {
		t.Fatal("stale purchase status must reject conditional acceptance")
	}
	var stored model.PurchaseRequest
	if err := db.First(&stored, p.ID).Error; err != nil {
		t.Fatal(err)
	}
	if stored.Status != "approved" || stored.DeviceID != 0 {
		t.Fatalf("stale purchase mutated: status=%s device=%d", stored.Status, stored.DeviceID)
	}
}
