package dto

import "testing"

func TestTransferEvidenceNormalizationKeepsInput(t *testing.T) {
	input := []string{" photo ", "", "certificate"}
	got := NormalizeTransferEvidence(input)
	if input[0] != " photo " || input[1] != "" {
		t.Fatalf("normalization mutated request: %q", input)
	}
	if len(got) != 2 || got[0] != "photo" {
		t.Fatalf("normalized evidence = %q", got)
	}
}
