package model

import (
	"testing"
	"time"
)

func TestMaintenanceScheduleUsesCalendarBoundaries(t *testing.T) {
	start := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
	next, err := NextMaintenanceDate(start, "monthly")
	if err != nil {
		t.Fatal(err)
	}
	want := time.Date(2024, 2, 1, 9, 0, 0, 0, time.UTC)
	if !next.Equal(want) {
		t.Fatalf("monthly date = %s, want %s", next, want)
	}
	if _, err := NextMaintenanceDate(start, "repair"); err == nil {
		t.Fatal("repair cannot be auto-scheduled")
	}
}
