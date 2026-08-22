package repository

import (
	"testing"

	"github.com/medasset/medasset/internal/model"
)

func TestScrapArchiveWriteDetachesMetadata(t *testing.T) {
	source := &model.ScrapRequest{Reason: "retired", ArchiveMetadata: map[string]string{"source": "approval"}}
	detached := detachScrapArchive(source)
	detached.ArchiveMetadata["source"] = "changed"
	if source.ArchiveMetadata["source"] != "approval" {
		t.Fatalf("archive write changed source metadata: %v", source.ArchiveMetadata)
	}
}
