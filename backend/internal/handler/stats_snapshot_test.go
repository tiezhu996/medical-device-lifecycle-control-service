package handler

import (
	"testing"

	"github.com/medasset/medasset/internal/dto"
)

func TestStatsResponseCapacityBoundary(t *testing.T) {
	groups := make([]dto.GroupStat, 2, 4)
	groups[0] = dto.GroupStat{Name: "A", Count: 1}
	groups[1] = dto.GroupStat{Name: "B", Count: 2}
	source := &dto.OverviewResp{Groups: groups}
	prepared := prepareStatsResponse(source)
	prepared.Groups = append(prepared.Groups, dto.GroupStat{Name: "C", Count: 3})
	if source.Groups[:3][2].Name == "C" {
		t.Fatal("handler response append reused source capacity")
	}
}
