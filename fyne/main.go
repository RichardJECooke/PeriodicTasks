package main

import (
	fileLogic "github.com/RichardJECooke/PeriodicTasks/src/5fileLogic"
	ui "github.com/RichardJECooke/PeriodicTasks/src/6ui"
)

func main() {
	fileLogic.Start()
	ui.Start()
}
