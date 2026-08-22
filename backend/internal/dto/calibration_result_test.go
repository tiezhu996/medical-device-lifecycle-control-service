package dto

import "testing"

func TestCalibrationRejectsIncompleteUnqualifiedResult(t *testing.T) {
	req := &CalibrationResultReq{Result: "unqualified"}
	if err := req.Validate(); err == nil {
		t.Fatal("unqualified result without evidence must be rejected")
	}
}
