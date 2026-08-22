package model

import "testing"

func TestAuditBatchClosesResultAfterWorkers(t *testing.T) {
	source := AuditLog{ID: 9, UserID: 4, Username: "nurse", Action: "PATCH", Module: "/devices", EntityID: "8", Detail: "status updated", IP: "127.0.0.1", RequestID: "request-42"}
	clone := source.CloneForBatch()
	if clone != source {
		t.Fatalf("batch copy lost audit fields: %#v", clone)
	}
}
