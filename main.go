// main.go
package main

import (
	"fmt"
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

	a := app.New()
	w := a.NewWindow("Unreal Project Launcher")
	view := NewAppView()
	view.SetProjects(catalog.Projects)

	var current *Project

	view.List.OnSelected = func(id int) {
		p := catalog.Projects[id]
		current = &p
		var engineName string
		if e := catalog.FindEngine(p.EngineID); e != nil {
			engineName = fmt.Sprintf("%s (%s)", e.Name, e.ID)
		} else {
			engineName = "Unknown Engine"
		}
		view.ShowDetails(p, engineName)
	}

	view.List.OnUnselected = func(id int) { current = nil }

	view.LaunchBtn.OnTapped = func() {
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
