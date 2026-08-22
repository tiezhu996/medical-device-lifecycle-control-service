package service

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/medasset/medasset/internal/model"
)

func TestAuditBatchWaitsForAllWriters(t *testing.T) {
	gate := make(chan struct{})
	logs := []model.AuditLog{{RequestID: "one"}, {RequestID: "two"}, {RequestID: "three"}}
	var writes atomic.Int32
	done := processAuditBatch(logs, func(*model.AuditLog) error {
		writes.Add(1)
		return nil
	}, gate)
	select {
	case <-done:
		t.Fatal("batch completed while every writer was still blocked")
	case <-time.After(50 * time.Millisecond):
	}
	close(gate)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("batch did not complete after releasing writers")
	}
	if writes.Load() != int32(len(logs)) {
		t.Fatalf("completed writes = %d, want %d", writes.Load(), len(logs))
	}
}
