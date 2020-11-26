package splashscreen

/*
	Модуль предоставляет шапку для вывода на экран
*/

import (
	"fmt"
)

// Print -- выводит шапку в начале работы
func Print(vers, build, data string) {
	fmt.Printf("\t\t/===========================\\\n")
	fmt.Printf("\t\t|    Оберон-компилятор      |\n")
	fmt.Printf("\t\t| Версия:%v Сборка:%v  |\n", vers, build)
	fmt.Printf("\t\t| Дата:%v  |\n", data)
	fmt.Printf("\t\t\\===========================/\n")
}
