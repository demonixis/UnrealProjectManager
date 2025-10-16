// model.go
package main

type Engine struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	EditorPath string `json:"editorPath"`
}

type Project struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Uproject string `json:"uproject"`
	EngineID string `json:"engineId"`
}

type Catalog struct {
	Engines  []Engine  `json:"engines"`
	Projects []Project `json:"projects"`
}

func (catalog Catalog) FindEngine(id string) *Engine {
	for i := range catalog.Engines {
		if catalog.Engines[i].ID == id {
			return &catalog.Engines[i]
		}
	}

	return nil
}

func (catalog Catalog) FindEngineIndex(id string) int {
	for i := range catalog.Engines {
		if catalog.Engines[i].ID == id {
			return i
		}
	}

	return -1
}
