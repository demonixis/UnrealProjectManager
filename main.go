// main.go
package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	jsonPath := "./catalog.json"
	if len(os.Args) > 1 {
		jsonPath = os.Args[1]
	}

	catalog, err := LoadCatalog(jsonPath)
	if err != nil {
		panic(err)
	}

	var engines []string
	for i := range catalog.Engines {
		engines = append(engines, catalog.Engines[i].ID)
	}

	a := app.New()
	w := a.NewWindow("Unreal Project Launcher")
	view := NewAppView(engines)
	view.SetProjects(catalog.Projects)

	var current *Project

	view.List.OnSelected = func(id int) {
		p := catalog.Projects[id]
		current = &p
		view.ShowDetails(p, catalog.FindEngineIndex(p.EngineID))
	}

	view.List.OnUnselected = func(id int) { current = nil }

	view.RunButton.OnTapped = func() {
		if current == nil {
			return
		}

		e := catalog.FindEngine(current.EngineID)
		if e != nil {
			return
		}
		_ = runEditor(e.EditorPath, current.Uproject)
	}

	w.SetContent(view.Root())
	w.Resize(fyne.NewSize(900, 560))
	w.ShowAndRun()
}

func runEditor(editorPath string, uproject string) error {
	return nil
}
