package main

import (
	"encoding/json"
	"os"
)

// parse JSON file (!! DisallowUnknownFields !!!
func loadJSON(path string) (*Root, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields() // be strict about schema

	var root Root
	if err := dec.Decode(&root); err != nil {
		return nil, err
	}

	// Basic sanity checks
	if root.Name == "" || len(root.Categories) == 0 {
		return nil, err
	}

	return &root, nil
}

// export JSON for import
func writeJSON(path string, obj *Root) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ") // pretty print with two-space indent
	enc.SetEscapeHTML(true) // keep \u0026 etc like your input; set to false to write &/< /> literally
	return enc.Encode(obj)  // writes trailing newline
}
