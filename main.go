package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	appInst := app.New()
	window := appInst.NewWindow("Hello")
	window.SetContent(widget.NewLabel("Hello Fyne!"))
	window.ShowAndRun()
}
