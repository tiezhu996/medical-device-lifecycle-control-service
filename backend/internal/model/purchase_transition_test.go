package model

import (
	"testing"
	"time"
)

func TestPurchaseTransitionRequiresDelivery(t *testing.T) {
	p := &PurchaseRequest{Status: "approved"}
	if p.CanAccept() {
		t.Fatal("approved purchase cannot skip delivery")
	}
	now := time.Now()
	p.Status = "delivered"
	p.DeliveredAt = &now
	if !p.CanAccept() {
		t.Fatal("delivered purchase without a device should be acceptable")
	}
	p.DeviceID = 8
	if p.CanAccept() {
		t.Fatal("purchase linked to a device cannot be accepted twice")
	}
}
