package dto

import "testing"

func TestPlanTypesAreUniqueAndSupported(t *testing.T) {
	types, err := DefaultMaintenancePlanTypes()
	if err != nil {
		t.Fatal(err)
	}
	if len(types) != 4 {
		t.Fatalf("plan types = %v", types)
	}
	seen := map[string]bool{}
	for _, value := range types {
		if seen[value] {
			t.Fatalf("duplicate plan type %q", value)
		}
		seen[value] = true
	}
}
