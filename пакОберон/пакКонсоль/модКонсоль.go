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

// Печатает любое сообщение
func (сам *ТуКонсоль) Печать(пСбщ string) {
	fmt.Println(пСбщ)
}

// Печатает сообщение об ошибке
func (сам *ТуКонсоль) Ошибка(пСбщ string) {
}
