package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadCatalog(path string) (Catalog, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return Catalog{}, fmt.Errorf("read json: %w", err)
	}

	var catalog Catalog
	if err := json.Unmarshal(fileContent, &catalog); err != nil {
		return Catalog{}, fmt.Errorf("parse json: %w", err)
	}

	return catalog, nil
}
