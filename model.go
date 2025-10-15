// model.go
package main

type Engine struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	EditorPath string `json:"editorPath"`
}

type Project struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Uproject string   `json:"uproject"`
	EngineID string   `json:"engineId"`
	Tags     []string `json:"tags,omitempty"`
	Notes    string   `json:"notes,omitempty"`
}

type Catalog struct {
	Engines  []Engine  `json:"engines"`
	Projects []Project `json:"projects"`
}

func (c Catalog) FindEngine(id string) *Engine {
	for i := range c.Engines {
		if c.Engines[i].ID == id {
			return &c.Engines[i]
		}
	}

	return nil
}
