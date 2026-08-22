package model

import "testing"

func TestTransferEvidenceSnapshotOwnsBackingArray(t *testing.T) {
	input := []string{"photo", "certificate"}
	snapshot := CopyTransferEvidence(input)
	snapshot[0] = "changed"
	if input[0] != "photo" {
		t.Fatalf("snapshot changed source: %q", input)
	}
}
