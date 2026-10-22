package main

import "testing"

func TestManifestRegistryConsistency(t *testing.T) {
	entries, err := loadComponentEntries()
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != len(componentFactories) {
		t.Fatalf("got %d manifest entries and %d factories", len(entries), len(componentFactories))
	}
	for _, entry := range entries {
		if entry.Build == nil {
			t.Errorf("manifest %q has no build function", entry.ID)
		}
		if _, ok := componentFactories[entry.ID]; !ok {
			t.Errorf("manifest %q has no matching registry factory", entry.ID)
		}
	}
}
