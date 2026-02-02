package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello, Fyne World")

	w.SetContent(widget.NewLabel("Hello, Fyne World"))
	w.ShowAndRun()
}
