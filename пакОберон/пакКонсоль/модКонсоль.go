package пакКонсоль

/*
	Модуль предоставляет средства культурного вывода в консоль различной информации
*/

import (
	мРес "../пакРесурс"
	мЛог "log"
)

//тКонсоль предоставляет средство для вывода различных классов информации
type тКонсоль struct {
}

var (
	//Конс -- универсальный объект (синглетон) для использования везде, где только можно
	Конс    *тКонсоль
	отладка = true
)

//Печать -- печатает любое сообщение
func (сам *тКонсоль) Печать(пСбщ string) {
	мЛог.Println(пСбщ)
}

//Отладить -- печатает отладочное сообщение
func (сам *тКонсоль) Отладить(пСбщ string) {
	if отладка == true {
		мЛог.Println(пСбщ + " ~")
	}
}

//Ошибка -- печатает сообщение об ошибке
func (сам *тКонсоль) Ошибка(пСбщ string) {
	мЛог.Println(пСбщ + " !!")
}

//ШапкаПечать -- печатает шапку в консоли
func (сам *тКонсоль) ШапкаПечать() {
	мЛог.Println()
	мЛог.Println("\t+-----------------------------+")
	мЛог.Println("\t|   Компилятор Oberon-07      |")
	мЛог.Println("\t|KBK Techniks ltd. " + мРес.Год + " BSD-2 |")
	мЛог.Println("\t|" + мРес.Дата + " " + мРес.Время + " Сборка " + мРес.Сборка + " |")
	мЛог.Println("\t+-----------------------------+")
	мЛог.Println()
}

func init() {
	Конс = &тКонсоль{}
}
