package пакОшибка

/*
Описывает струтктуру, в которой хранится ошибка возникшая в процессе-хозяине.
Ошибки делятся на два вида:
1. Внутренняя
2. Исходника.
*/

import (
	мКонс "../../../пакКонсоль"
	"os"
)

//Синоним используемый только для ошибок
type тОш bool

//ТОшибка -- структура для хранения внутренних ошибок
type ТОшибка struct {
	бВнутр тОш    // Ошибка самого компилятора
	бИсх   тОш    // Ошибка исходника
	стрТип string // Тип который вызывает ошибку
}

//ОшибкаНов -- возвращает ссылку на ТОшибка
func ОшибкаНов(стрТип string) (ош *ТОшибка) {
	ош = &ТОшибка{
		бВнутр: false,
		бИсх:   false,
		стрТип: стрТип,
	}

	return ош
}

//Внутр -- внутренняя ошибка компилятора
func (сам *ТОшибка) Внутр(пВызов string, пОш string) {
	сам.бВнутр = true
	стрСообщ := сам.стрТип + "." + пВызов + "(): ошибка компилятора." + пОш
	мКонс.Конс.Ошибка(стрСообщ)
	os.Exit(1)
}

//Исх -- ошибка исходника
func (сам *ТОшибка) Исх(пВызов string, пОш string) {
	сам.бИсх = true
	мКонс.Конс.Ошибка(сам.стрТип + "." + пВызов + "(): ошибка исходника. " + пОш)
}
