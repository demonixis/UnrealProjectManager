package main

import (
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
	Projects    []Project
}

func NewAppView() *AppView {
	v := &AppView{}

	title := widget.NewLabel("Project Details")
	title.TextStyle = fyne.TextStyle{Bold: true}

	v.EngineLabel = widget.NewLabel("")
	v.LaunchBtn = widget.NewButtonWithIcon("Launch Editor", theme.ConfirmIcon(), nil)
	v.OpenFolder = widget.NewButtonWithIcon("Open Folder", theme.ConfirmIcon(), nil)
	v.OpenUProj = widget.NewButtonWithIcon("Open Project", theme.ConfirmIcon(), nil)

	btns := container.New(layout.NewGridLayout(3), v.LaunchBtn, v.OpenFolder, v.OpenUProj)
	right := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewRichTextFromMarkdown("**Engine**"),
		v.EngineLabel,
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
			project := v.Projects[i]
			container := o.(*fyne.Container)
			projectIcon := GetProjectIcon(project.Uproject, fyne.NewSize(48, 48))
			container.Objects[0] = projectIcon
			projectNameLabel := container.Objects[1].(*widget.Label)
			projectNameLabel.SetText(project.Name)
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
}
