// Package gui -- главный тип графической оболочки компилятора
package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/prospero78/goOC/internal/gui/winabout"
	"github.com/prospero78/goOC/internal/gui/winmain"
)

// TGui -- операции сграфической оболочкой
type TGui struct {
	app      fyne.App
	winMain  *winmain.TWinMain
	winAbout *winabout.TWinAbout
}

// New -- возвращает новый *TGui
func New() *TGui {
	return &TGui{}
}

// Run -- запускает графическую оболочку в работу
func (sf *TGui) Run() {
	sf.app = app.New()
	sf.winAbout = winabout.New()
	sf.winMain = winmain.New(sf.app, sf.winAbout.Show)
	sf.winMain.Show()
	sf.app.Run()
}
