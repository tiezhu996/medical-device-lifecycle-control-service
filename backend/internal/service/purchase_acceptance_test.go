package service

import (
	"testing"
	"time"

	"github.com/medasset/medasset/internal/dto"
	"github.com/medasset/medasset/internal/model"
	"github.com/medasset/medasset/internal/repository"
)

func TestAcceptanceRollsBackDeviceWhenPurchaseUpdateFails(t *testing.T) {
	env := newTestServiceEnv(t)
	now := time.Now()
	p := &model.PurchaseRequest{
		RequestNo: "PR-ROLLBACK", Department: "ICU", ApplicantName: "Li",
		DeviceName: "Monitor", Quantity: 1, Status: "delivered", DeliveredAt: &now,
	}
	if err := env.db.Create(p).Error; err != nil {
		t.Fatal(err)
	}
	if err := env.db.Exec(`CREATE TRIGGER reject_acceptance BEFORE UPDATE ON purchase_requests
		WHEN NEW.status = 'accepted' BEGIN SELECT RAISE(ABORT, 'acceptance write failed'); END;`).Error; err != nil {
		t.Fatal(err)
	}
	svc := NewPurchaseService(repository.NewPurchaseRepository(env.db), repository.NewDeviceRepository(env.db), env.audit, env.logger)
	_, err := svc.Accept(p.ID, &dto.AcceptReq{
		AcceptancePerson: "engineer", AssetCode: "MD-ROLLBACK", CertificateNo: "CERT-1",
		RegistrationNo: "REG-1", PartsList: "probe", WarrantyMonths: 12,
	}, "admin")
	if err == nil {
		t.Fatal("triggered purchase update failure must be returned")
	}
	var devices int64
	if err := env.db.Model(&model.Device{}).Count(&devices).Error; err != nil {
		t.Fatal(err)
	}
	if devices != 0 {
		t.Fatalf("device creation escaped rollback: %d", devices)
	}
}
