package service

import (
	"fmt"
	"sync"
	"testing"

	"github.com/medasset/medasset/internal/model"
)

func TestDeviceDepartmentAggregationHasNoRace(t *testing.T) {
	var cache deviceSnapshotCache
	cache.Store(makeDevices("initial"))
	oldSnapshot := cache.Load()
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-start
		for i := 0; i < 2000; i++ {
			cache.Store(makeDevices(fmt.Sprintf("refresh-%d", i)))
		}
	}()
	go func() {
		defer wg.Done()
		<-start
		for i := 0; i < 2000; i++ {
			_ = oldSnapshot[i%len(oldSnapshot)].Department
		}
	}()
	close(start)
	wg.Wait()
	if oldSnapshot[0].Department != "initial" {
		t.Fatalf("old snapshot was overwritten: %q", oldSnapshot[0].Department)
	}
}

func makeDevices(department string) []model.Device {
	devices := make([]model.Device, 32)
	for i := range devices {
		devices[i] = model.Device{ID: uint(i + 1), Department: department}
	}
	return devices
}
