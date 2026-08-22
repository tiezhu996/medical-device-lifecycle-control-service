package repository

import (
	"testing"

	"github.com/medasset/medasset/internal/model"
)

func TestTransferRepositoryDetachesEvidenceBeforeWrite(t *testing.T) {
	source := &model.TransferRequest{Evidence: []string{"photo", "certificate"}}
	detached := detachTransfer(source)
	detached.Evidence[0] = "changed"
	if source.Evidence[0] != "photo" {
		t.Fatalf("repository write copy changed source: %q", source.Evidence)
	}
}
