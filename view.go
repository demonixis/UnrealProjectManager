package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AppView struct {
	root        fyne.CanvasObject
	List        *widget.List
	LaunchBtn   *widget.Button
	OpenFolder  *widget.Button
	OpenUProj   *widget.Button
	EngineLabel *widget.Label
	Notes       *widget.Label
	Projects    []Project
}

func NewAppView() *AppView {
	v := &AppView{}

	title := widget.NewLabel("Select a project")
	title.TextStyle = fyne.TextStyle{Bold: true}

	v.EngineLabel = widget.NewLabel("")
	v.Notes = widget.NewLabel("")
	v.Notes.Wrapping = fyne.TextWrapWord

	v.LaunchBtn = widget.NewButtonWithIcon("Launch Editor", theme.ConfirmIcon(), nil)
	v.OpenFolder = widget.NewButtonWithIcon("Open Folder", theme.ConfirmIcon(), nil)
	v.OpenUProj = widget.NewButtonWithIcon("Open .uproject", theme.ConfirmIcon(), nil)

	btns := container.New(layout.NewGridLayout(3), v.LaunchBtn, v.OpenFolder, v.OpenUProj)
	right := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewRichTextFromMarkdown("**Engine"),
		v.EngineLabel,
		widget.NewRichTextFromMarkdown("**Notes"),
		v.Notes,
		widget.NewSeparator(),
		btns,
	)

	v.List = widget.NewList(
		func() int { return len(v.Projects) },
		func() fyne.CanvasObject {
			return container.New(layout.NewHBoxLayout(),
				widget.NewIcon(theme.FileApplicationIcon()),
				widget.NewLabel(""),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)
			lbl := c.Objects[1].(*widget.Label)
			p := v.Projects[i]
			lbl.SetText(fmt.Sprintf("%s [%s]", p.Name, strings.Join(p.Tags, ", ")))
		},
	)

	split := container.NewHSplit(v.List, container.NewPadded(right))
	split.SetOffset(0.65)
	v.root = split

	return v
}

func (v *AppView) Root() fyne.CanvasObject { return v.root }

func (v *AppView) SetProjects(ps []Project) {
	v.Projects = append([]Project(nil), ps...)
	v.List.Refresh()
}

func (v *AppView) ShowDetails(p Project, engineName string) {
	v.EngineLabel.SetText(engineName)
	if strings.TrimSpace(p.Notes) == "" {
		v.Notes.SetText("(no notes)")
	} else {
		v.Notes.SetText(p.Notes)
	}
}
