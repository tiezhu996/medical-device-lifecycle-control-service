package service

import (
	"testing"

	"github.com/medasset/medasset/internal/model"
	"github.com/medasset/medasset/internal/repository"
)

func TestGeneratePlansRollsBackFailedDeviceBatch(t *testing.T) {
	env := newTestServiceEnv(t)
	device := &model.Device{AssetCode: "PLAN-1", Barcode: "B-PLAN-1", Name: "Ventilator", Status: "in_use"}
	if err := env.db.Create(device).Error; err != nil {
		t.Fatal(err)
	}
	if err := env.db.Exec(`CREATE TRIGGER reject_plan BEFORE INSERT ON maintenance_records
		BEGIN SELECT RAISE(ABORT, 'plan storage unavailable'); END;`).Error; err != nil {
		t.Fatal(err)
	}
	svc := NewMaintenanceService(repository.NewMaintenanceRepository(env.db), repository.NewDeviceRepository(env.db), env.audit, env.logger)
	created, err := svc.GeneratePlans("scheduler")
	if err == nil {
		t.Fatal("batch storage failure must be returned")
	}
	if created != 0 {
		t.Fatalf("reported %d uncommitted plans", created)
	}
	var stored int64
	if err := env.db.Model(&model.MaintenanceRecord{}).Count(&stored).Error; err != nil {
		t.Fatal(err)
	}
	if stored != 0 {
		t.Fatalf("failed device batch left %d records", stored)
	}
}
