package main

import (
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := fyneApp.New()
	window := app.NewWindow("Hello World")

	window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
