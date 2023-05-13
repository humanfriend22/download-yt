package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var w fyne.Window

func Throw(err error) {
	dialog.ShowError(err, w)
}
