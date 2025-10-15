package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func removeDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("Error removing at %s", path)
	}
}

func GetProjectIcon(uproject string, size fyne.Size) fyne.CanvasObject {
	iconPath := strings.TrimSuffix(uproject, filepath.Ext(uproject)) + ".png"

	if _, err := os.Stat(iconPath); err == nil {
		img := canvas.NewImageFromFile(iconPath)
		img.SetMinSize(size)
		img.FillMode = canvas.ImageFillContain
		return img
	}

	return widget.NewIcon(theme.FileIcon())
}

func GetProjectIconName(uproject string) string {
	file := filepath.Base(uproject)
	name := file[:len(file)-len(filepath.Ext(file))]
	return name + ".png"
}

func GetProjectPath(uproject string) string {
	return filepath.Dir(uproject)
}

func GetProjectName(uproject string) string {
	path := filepath.Dir(uproject)
	base := filepath.Base(path)
	return base[:len(base)-len(filepath.Ext(base))]
}

func CleanUnrealProject(project Project) {
	dirs := []string{
		"Binaries",
		"DerivedDataCache",
		"Intermediate",
		"Saved",
		"Script",
	}

	path := GetProjectPath(project.Uproject)

	for _, dir := range dirs {
		fullPath := filepath.Join(path, dir)
		removeDir(fullPath)
	}

	pluginsPath := filepath.Join(path, "Plugins")
	entries, err := os.ReadDir(pluginsPath)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pluginDir := filepath.Join(pluginsPath, entry.Name())
		removeDir(filepath.Join(pluginDir, "Binaries"))
		removeDir(filepath.Join(pluginDir, "Intermediate"))
	}
}

func GenerateUnrealSolution(uproject string, uePath string) *exec.Cmd {
	switch runtime.GOOS {
	case "linux":
		buildCmd := uePath + "/Engine/Build/BatchFiles/Linux/GenerateProjectFiles.sh"
		return exec.Command(buildCmd, uproject, "-game")

	case "darwin":
		buildCmd := uePath + "/Engine/Build/BatchFiles/Mac/GenerateProjectFiles.sh"
		return exec.Command(buildCmd, uproject, "-game")

	case "windows":
		ubtPath := uePath + "\\Engine\\Binaries\\DotNET\\UnrealBuildTool\\UnrealBuildTool.dll"
		return exec.Command(
			"dotnet", ubtPath,
			"-projectfiles",
			"-project="+uproject,
			"-game",
			"-progress",
		)
	}

	return nil
}

func BuildUnrealSolution(uproject string, uePath string) *exec.Cmd {
	projectPath := GetProjectPath(uproject)
	projectName := GetProjectName(uproject)

	switch runtime.GOOS {
	case "linux":
		return exec.Command("make", "-C", projectPath, projectName)

	case "windows":
		buildBatch := uePath + "\\Engine\\Build\\BatchFiles\\Build.bat"
		target := projectName + "Editor"
		platform := "Win64"
		config := "Development"

		return exec.Command(
			buildBatch,
			target,
			platform,
			config,
			"-Project="+uproject,
			"-WaitMutex",
		)
	}

	return nil
}

func RunUnrealProject(uproject string, uePath string) *exec.Cmd {
	switch runtime.GOOS {
	case "linux":
		return exec.Command(uePath+"/Engine/Binaries/Linux/UnrealEditor", uproject)
	case "darwin":
		return exec.Command(uePath+"/Engine/Binaries/Mac/UnrealEditor.app/Contents/MacOS/UnrealEditor", uproject)
	case "windows":
		return exec.Command(uePath+"/Engine/Binaries/Win64/UnrealEditor.exe", uproject)
	}

	return nil
}
