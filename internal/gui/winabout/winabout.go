// Package winabout -- показывает окно о программе
package winabout

import (
	"fyne.io/fyne/v2"
)

// TWinAbout -- операции с окном "О программе"
type TWinAbout struct {
	winAbout fyne.Window
}

// New -- возвращает новый *TWinAbout
func New() *TWinAbout {
	wa := &TWinAbout{
		winAbout: fyne.CurrentApp().NewWindow("О программе"),
	}
	wa.winAbout.Resize(fyne.Size{Width: 320, Height: 240})
	return wa
}

// Show -- показывает окно "О программе"
func (sf *TWinAbout) Show() {
	sf.winAbout.Show()
}
