package service

import "testing"

func TestTransferApprovalDoesNotMutateSubmittedEvidence(t *testing.T) {
	input := make([]string, 2, 4)
	input[0], input[1] = " photo ", "certificate"
	got := appendApprovalEvidence(input, "reviewer")
	if input[0] != " photo " || input[1] != "certificate" {
		t.Fatalf("approval mutated submitted evidence: %q", input)
	}
	if input[:3][2] != "" {
		t.Fatalf("approval reused submitted capacity: %q", input[:3])
	}
	if len(got) != 3 || got[2] != "approved:reviewer" {
		t.Fatalf("approval evidence = %q", got)
	}
}
