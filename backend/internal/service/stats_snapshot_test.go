package service

import "testing"

func TestStatsOverviewKeepsPreviousSnapshot(t *testing.T) {
	var buffer statsGroupBuffer
	first := buffer.Build(map[string]int64{"ICU": 2, "ER": 5})
	want := map[string]int64{}
	for _, item := range first {
		want[item.Name] = item.Count
	}
	buffer.Build(map[string]int64{"Lab": 9, "Ward": 11})
	for _, item := range first {
		if want[item.Name] != item.Count {
			t.Fatalf("previous overview drifted after refresh: %#v", first)
		}
	}
}
