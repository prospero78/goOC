// Package winmain -- главное окно для компилятора
package winmain

import (
	"fyne.io/fyne/v2"
	"github.com/sirupsen/logrus"

	"github.com/prospero78/goOC/internal/gui/winmain/frmmenu"
)

// TWinMain -- операции с главным окном
type TWinMain struct {
	win  fyne.Window
	menu *frmmenu.TFrmMenu
}

// New -- возвращает новый *TWinMain
func New(root fyne.App, fnAboutShow func()) *TWinMain {
	{ // Предусловия
		if root == nil {
			logrus.Panicln("winmain.go/New(): root==nil")
		}
		if fnAboutShow == nil {
			logrus.Panicln("winmain.go/New(): fnAboutShow==nil")
		}
	}
	wm := &TWinMain{
		win: root.NewWindow("goOC compiler"),
	}
	wm.win.Resize(fyne.Size{Width: 320, Height: 240})
	wm.menu = frmmenu.New(wm.win, fnAboutShow)
	return wm
}

// Show -- показать главное окно
func (sf *TWinMain) Show() {
	sf.win.Show()
}
