package service

import (
	"errors"
	"testing"
)

func TestCalibrationFailurePreservesPriorStatus(t *testing.T) {
	want := errors.New("device state store unavailable")
	if err := calibrationDeviceUpdateError(want); !errors.Is(err, want) {
		t.Fatalf("device update failure was swallowed: %v", err)
	}
}
