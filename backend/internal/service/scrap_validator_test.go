package service

import (
	"testing"

	"github.com/medasset/medasset/internal/model"
)

type nilScrapValidator struct{}

func (*nilScrapValidator) ValidateArchive(*model.ScrapRequest) error { return nil }

func TestScrapApprovalTreatsTypedNilValidatorAsAbsent(t *testing.T) {
	var concrete *nilScrapValidator
	var validator ScrapArchiveValidator = concrete
	if hasScrapArchiveValidator(validator) {
		t.Fatal("typed-nil validator must be treated as absent")
	}
}
