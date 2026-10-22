package components

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

var validManifestStatuses = map[string]bool{"draft": true, "experimental": true, "stable": true, "deprecated": true}
var validManifestPropertyTypes = map[string]bool{"string": true, "float32": true, "bool": true, "int": true}

func TestManifestsAreValid(t *testing.T) {
	paths, err := filepath.Glob("*.manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		t.Fatal("no component manifests found")
	}
	seen := make(map[string]string)
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var manifest ComponentManifest
		if err := json.Unmarshal(data, &manifest); err != nil {
			t.Fatalf("%s: invalid JSON: %v", path, err)
		}
		if manifest.ID == "" || manifest.Name == "" || manifest.Description == "" {
			t.Errorf("%s: id, name, and description are required", path)
		}
		if previous, ok := seen[manifest.ID]; ok {
			t.Errorf("duplicate component ID %q in %s and %s", manifest.ID, previous, path)
		}
		seen[manifest.ID] = path
		if !validManifestStatuses[manifest.Status] {
			t.Errorf("%s: invalid status %q", path, manifest.Status)
		}
		if manifest.Variants != nil && len(manifest.Variants) == 0 {
			t.Errorf("%s: variants must not be empty when declared", path)
		}
		for _, property := range manifest.Properties {
			if property.Name == "" || !validManifestPropertyTypes[property.Type] {
				t.Errorf("%s: invalid property %q of type %q", path, property.Name, property.Type)
			}
		}
	}
}
