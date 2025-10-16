package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AppView struct {
	root           fyne.CanvasObject
	List           *widget.List
	LogText        *widget.Entry
	RunButton      *widget.Button
	CleanBtn       *widget.Button
	GenerateBtn    *widget.Button
	BuildBtn       *widget.Button
	EngineChoice   *widget.Select
	PlatformChoice *widget.Select
	BuildPlatform  *widget.Button
	Projects       []Project
}

func NewAppView(engines []string) *AppView {
	appView := &AppView{}

	title := widget.NewLabel("Project Details")
	title.TextStyle = fyne.TextStyle{Bold: true}

	appView.LogText = widget.NewMultiLineEntry()
	appView.LogText.SetMinRowsVisible(10)
	appView.EngineChoice = widget.NewSelect(engines, nil)
	appView.CleanBtn = widget.NewButtonWithIcon("Clean", theme.CancelIcon(), nil)
	appView.GenerateBtn = widget.NewButtonWithIcon("Generate Solution", theme.HistoryIcon(), nil)
	appView.BuildBtn = widget.NewButtonWithIcon("Build Sources", theme.ComputerIcon(), nil)
	appView.RunButton = widget.NewButtonWithIcon("Run Editor", theme.ConfirmIcon(), nil)
	appView.BuildPlatform = widget.NewButtonWithIcon("Build Platform", theme.GridIcon(), nil)
	appView.PlatformChoice = widget.NewSelect([]string{"Windows", "Linux", "Mac", "Android", "iOS", "VisionOS"}, nil)

	engineActs := container.New(
		layout.NewGridLayout(3),
		widget.NewLabel("Engine"),
		appView.EngineChoice,
		appView.RunButton,
	)

	btns := container.New(
		layout.NewGridLayout(3),
		appView.CleanBtn,
		appView.GenerateBtn,
		appView.BuildBtn)

	acts := container.New(
		layout.NewGridLayout(3),
		widget.NewLabel("Platform"),
		appView.PlatformChoice,
		appView.BuildPlatform,
	)

	right := container.New(
		layout.NewVBoxLayout(),
		title,
		widget.NewSeparator(),
		engineActs,
		layout.NewSpacer(),
		widget.NewSeparator(),
		widget.NewRichTextFromMarkdown("**Source Management**"),
		btns,
		widget.NewSeparator(),
		widget.NewRichTextFromMarkdown("**Build Settings**"),
		acts,
	)

	appView.List = widget.NewList(
		func() int { return len(appView.Projects) },
		func() fyne.CanvasObject {
			return container.New(layout.NewHBoxLayout(),
				widget.NewIcon(theme.FileApplicationIcon()),
				widget.NewLabel(""),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			project := appView.Projects[i]
			container := o.(*fyne.Container)
			projectIcon := GetProjectIcon(project.Uproject, fyne.NewSize(96, 96))
			container.Objects[0] = projectIcon
			projectNameLabel := container.Objects[1].(*widget.Label)
			projectNameLabel.SetText(project.Name)
		},
	)

	split := container.NewHSplit(appView.List, container.NewPadded(right))
	split.SetOffset(0.65)

	mainContainer := container.NewBorder(nil, appView.LogText, nil, nil, split)
	appView.root = mainContainer

	return appView
}

func (v *AppView) Root() fyne.CanvasObject { return v.root }

func (v *AppView) SetProjects(ps []Project) {
	v.Projects = append([]Project(nil), ps...)
	v.List.Refresh()
}

func (v *AppView) ShowDetails(p Project, engineId int) {
	v.EngineChoice.SetSelectedIndex(engineId)
}
