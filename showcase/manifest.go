package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/yswa-var/fynefactory/components"
)

func componentManifestDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "components")
}

func loadComponentManifests() ([]components.ComponentManifest, error) {
	paths, err := filepath.Glob(filepath.Join(componentManifestDir(), "*.manifest.json"))
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	manifests := make([]components.ComponentManifest, 0, len(paths))
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		var manifest components.ComponentManifest
		if err := json.Unmarshal(data, &manifest); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		manifests = append(manifests, manifest)
	}
	return manifests, nil
}

func loadComponentEntries() ([]ComponentEntry, error) {
	manifests, err := loadComponentManifests()
	if err != nil {
		return nil, err
	}
	entries := make([]ComponentEntry, 0, len(manifests))
	for _, manifest := range manifests {
		build, ok := componentFactories[manifest.ID]
		if !ok {
			return nil, fmt.Errorf("manifest %q has no showcase factory", manifest.ID)
		}
		entries = append(entries, ComponentEntry{
			ID: manifest.ID, Name: manifest.Name, Category: manifest.Category,
			Description: manifest.Description, Status: manifest.Status,
			Variants: manifest.Variants, Properties: manifest.Properties, Build: build,
		})
	}
	return entries, nil
}
