// Package winmain -- главное окно для компилятора
package winmain

import (
	"fyne.io/fyne/v2"
)

// TWinMain -- операции с главным окном
type TWinMain struct {
	win fyne.Window
}

// New -- возвращает новый *TWinMain
func New(root fyne.App) *TWinMain {
	wm := &TWinMain{
		win: root.NewWindow("goOC compiler"),
	}
	wm.win.Resize(fyne.Size{Width: 320, Height: 240})
	return wm
}

// Show -- показать главное окно
func (sf *TWinMain) Show() {
	sf.win.Show()
}
