// Package frmmenu -- верхний фрейм с меню
package frmmenu

import "fyne.io/fyne/v2"

// TFrmMenu -- операции с меню приложения
type TFrmMenu struct {
	mnuAbout fyne.Widget
}

// New -- возвращает новый *TFrmMenu
func New() *TFrmMenu {
	return &TFrmMenu{}
}
