package main

import (
	"time"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := fyneApp.New()
	window := app.NewWindow("Hello World")
	message := widget.NewLabel("Welcome")
	button := widget.NewButton("Update", func() {
		formatted := time.Now().Format("Time: 03:04:05")
		message.SetText(formatted)
	})
	window.SetContent(container.NewVBox(message, button))
	window.Resize(fyne.NewSize(1280, 720))
	window.CenterOnScreen()
	window.ShowAndRun()
}
