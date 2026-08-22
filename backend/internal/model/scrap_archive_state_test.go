package model

import "testing"

func TestScrapArchiveMetadataPreservesLegacyReason(t *testing.T) {
	metadata := EnsureScrapArchiveMetadata(nil, "beyond repair")
	if metadata == nil || metadata["reason"] != "beyond repair" {
		t.Fatalf("legacy reason missing: %v", metadata)
	}
}
