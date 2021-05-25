// Package frmmenu -- верхний фрейм с меню
package frmmenu

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// TFrmMenu -- операции с меню приложения
type TFrmMenu struct {
	mnuMain     *fyne.MainMenu
	mnuAbout    widget.Menu
	fnShowAbout func()
}

// New -- возвращает новый *TFrmMenu
func New(win fyne.Window, fnAboutShow func()) *TFrmMenu {
	{ // Предусловия
		if win == nil {
			logrus.Panicln("winmain.go/New(): winMain==nil")
		}
		if fnAboutShow == nil {
			logrus.Panicln("winmain.go/New(): fnAboutShow==nil")
		}
	}
	menu := &TFrmMenu{
		fnShowAbout: fnAboutShow,
	}
	itemAbout := fyne.NewMenuItem("О програме", menu.showAbout)
	mnuAbout := fyne.NewMenu("Помощь", itemAbout)
	mnuMain := fyne.NewMainMenu(mnuAbout)
	menu.mnuMain = mnuMain
	win.SetMainMenu(mnuMain)
	return menu
}

// Показывает окно "О программе"
func (sf *TFrmMenu) showAbout() {
	fmt.Printf("TFmMenu.showAbout()\n")
	sf.fnShowAbout()
}
