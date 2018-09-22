// модКонсоль
package пакКонсоль

/*
	Предоставляет средства культурного вывода в консоль различной информации
*/

import (
	"fmt"
)

type ТуКонсоль struct {
}

var (
	Конс    *ТуКонсоль
	отладка = true
)

// Печатает любое сообщение
func (сам *ТуКонсоль) Печать(пСбщ string) {
	fmt.Println(пСбщ)
}

// Печатает отладочное сообщение
func (сам *ТуКонсоль) Отладить(пСбщ string) {
	if отладка == true {
		fmt.Println(пСбщ)
	}
}

// Печатает сообщение об ошибке
func (сам *ТуКонсоль) Ошибка(пСбщ string) {
}

func init() {
	Конс = new(ТуКонсоль)
}
