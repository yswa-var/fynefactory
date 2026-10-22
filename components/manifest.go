package components

type ComponentManifest struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	Variants    []string           `json:"variants,omitempty"`
	Properties  []ManifestProperty `json:"properties"`
}

type ManifestProperty struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Default any    `json:"default"`
}
