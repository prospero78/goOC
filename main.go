// main
package main

import (
	пакОбер "./пакОберон"
	"fmt"
)

func main() {
	var оберон = пакОбер.Новый()
	оберон.Компилер = 0
	оберон.Выполнить()
	fmt.Println("Привет, мир!")
}
