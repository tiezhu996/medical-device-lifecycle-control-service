package dto

import "testing"

func TestScrapApprovalInitializesArchiveMetadata(t *testing.T) {
	metadata := (&ScrapApproveReq{}).ArchiveMetadata()
	metadata["operator"] = "reviewer"
	if metadata["source"] != "scrap_approval" {
		t.Fatalf("archive source missing: %v", metadata)
	}
}
