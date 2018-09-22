// модОшибка
package пакОшибка

/*
Описывает струтктуру, в которой хранится ошибка возникшая в процессе-хозяине.
Ошибки делятся на два вида:
1. Внутренняя
2. Исходника.
*/

import (
	пакКонс "../../../пакКонсоль"
)

//Синоним используемый только для ошибок
type тОш bool

type ТуОшибка struct {
	бВнутр тОш                // Ошибка самого компилятора
	бИсх   тОш                // Ошибка исходника
	стрТип string             // Тип который вызывает ошибку
	конс   *пакКонс.ТуКонсоль // Ссылка на консоль
}

func Новый(стрТип string, пКонс *пакКонс.ТуКонсоль) (ош *ТуОшибка) {
	ош = new(ТуОшибка)
	ош.бВнутр = false
	ош.бИсх = false
	ош.стрТип = стрТип
	ош.конс = пКонс
	return ош
}

func (сам *ТуОшибка) Внутр(пВызов string, пОш string) {
	сам.бВнутр = true
	сам.конс.Ошибка(сам.стрТип + "." + пВызов + "(): ошибка компилятора." + пОш)
}

func (сам *ТуОшибка) Исх(пВызов string, пОш string) {
	сам.бИсх = true
	сам.конс.Ошибка(сам.стрТип + "." + пВызов + "(): ошибка исходника. " + пОш)
}

func (сам *ТуОшибка) БВнутр() тОш {
	return сам.бВнутр
}

func (сам *ТуОшибка) БИсх() тОш {
	return сам.бИсх
}
