package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadCatalog(path string) (Catalog, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Catalog{}, fmt.Errorf("read json: %w", err)
	}
	var c Catalog
	if err := json.Unmarshal(b, &c); err != nil {
		return Catalog{}, fmt.Errorf("parse json: %w", err)
	}

	return c, nil
}
