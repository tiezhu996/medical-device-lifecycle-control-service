package model

import "testing"

func TestCalibrationResultStateBlocksRepeatedFailure(t *testing.T) {
	record := &CalibrationRecord{Status: "unqualified", CalibrationCycleMonths: 12}
	if record.CanRecordResult() {
		t.Fatal("unqualified terminal record cannot accept another result")
	}
	record.Status = "due"
	if !record.CanRecordResult() {
		t.Fatal("due record with a valid cycle should accept a result")
	}
}
