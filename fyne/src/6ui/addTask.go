package ui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/RichardJECooke/PeriodicTasks/src"
)

func CreateAddTask() *widget.Button {
	result := widget.NewButton("Add task", func() {
		src.AddTask()
	})
	return result
}
