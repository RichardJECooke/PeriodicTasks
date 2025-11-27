package ui

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Start() {
	app := fyneApp.New()
	window := app.NewWindow("Periodic Tasks")
	message := widget.NewLabel("Tasks")
	appContainer := container.NewVBox(
		message,
		container.NewHBox(CreateAddTask()),
		CreateTaskList(),
	)
	window.SetContent(appContainer)
	window.Resize(fyne.NewSize(1280, 720))
	window.CenterOnScreen()
	window.ShowAndRun()
}
