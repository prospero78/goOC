// Package winabout -- показывает окно о программе
package winabout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// TWinAbout -- операции с окном "О программе"
type TWinAbout struct {
	winAbout fyne.Window
	btnHide  *widget.Button
}

// New -- возвращает новый *TWinAbout
func New() *TWinAbout {
	wa := &TWinAbout{
		winAbout: fyne.CurrentApp().NewWindow("О программе"),
	}
	wa.winAbout.Resize(fyne.Size{Width: 320, Height: 240})
	txtAbout := widget.NewMultiLineEntry()
	txtAbout.Text = `Программа создана в рамках проекта по созданию Оберон-компилятора
на golang.

2021 (c) prospero78su

Лицензия: BSD-2`
	boxTop := container.New(layout.NewMaxLayout(), txtAbout)
	wa.btnHide = widget.NewButton("Закрыть", wa.Hide)
	layBottom := layout.NewBorderLayout(nil, wa.btnHide, nil, nil)
	boxBottom := container.New(layBottom, wa.btnHide)
	boxPack := container.New(layout.NewBorderLayout(boxTop, boxBottom, nil, nil), boxTop, boxBottom)
	wa.winAbout.SetContent(boxPack)
	return wa
}

// Hide -- скрывает окно "О программе"
func (sf *TWinAbout) Hide() {
	sf.winAbout.Hide()
}

// Show -- показывает окно "О программе"
func (sf *TWinAbout) Show() {
	sf.winAbout.Show()
}

// Close -- типа закрывает окно
func (sf *TWinAbout) Close() {
	sf.winAbout.Close()
}
