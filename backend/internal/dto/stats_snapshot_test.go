package dto

import "testing"

func TestStatsGroupsOwnStorage(t *testing.T) {
	source := &OverviewResp{Groups: []GroupStat{{Name: "ICU", Count: 4}, {Name: "ER", Count: 7}}}
	clone := CloneOverview(source)
	clone.Groups[0].Count = 99
	if source.Groups[0].Count != 4 {
		t.Fatalf("response clone changed source groups: %#v", source.Groups)
	}
}
