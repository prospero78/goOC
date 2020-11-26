package splashscreen

/*
	Модуль предоставляет шапку для вывода на экран
*/

import (
	"fmt"
	"oc/internal/app/param"
)

// Print -- выводит шапку в начале работы
func Print(param *param.TParam) {
	fmt.Printf("\t\t/===========================\\\n")
	fmt.Printf("\t\t|    Оберон-компилятор      |\n")
	fmt.Printf("\t\t| Версия:%v Сборка:%v  |\n", param.Vers, param.Build)
	fmt.Printf("\t\t| Дата:%v  |\n", param.Data)
	fmt.Printf("\t\t\\===========================/\n")
}
