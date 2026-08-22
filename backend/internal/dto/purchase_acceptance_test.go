package dto

import "testing"

func TestAcceptancePayloadRequiresEvidence(t *testing.T) {
	req := &AcceptReq{AcceptancePerson: "engineer", AssetCode: "MD-101"}
	if err := req.Validate(); err == nil {
		t.Fatal("acceptance without certificate evidence must be rejected")
	}
}
