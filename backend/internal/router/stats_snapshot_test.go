package router

import (
	"testing"

	"github.com/medasset/medasset/internal/dto"
)

func TestStatsRouteReturnsStablePayload(t *testing.T) {
	groups := []dto.GroupStat{{Name: "Radiology", Count: 6}}
	payload := routeStatsSnapshot(groups)
	groups[0].Count = 40
	if payload[0].Count != 6 {
		t.Fatalf("routed payload changed after source reuse: %#v", payload)
	}
}
