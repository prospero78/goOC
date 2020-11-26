package main

/*
	Модуль определяет точку входа всего приложения
*/

import (
	"oc/internal/app"
	"oc/internal/app/param"
	"oc/internal/log"
	mSs "oc/internal/splashscreen"
)

func main() {
	режим := log.КОтладка
	парам := &param.ТПарам{
		Версия: "0.0.6",
		Сборка: "0070",
		Дата:   "2020-08-31 22:03:00",
		Режим:  режим,
	}
	mSs.Печать(парам)
	лог := log.Нов("main", режим)
	лог.Отладка("main", "Начало работы компилятора")
	прилож, ош := app.Нов(парам)
	if ош != nil {
		лог.Паника("main", "Запуск приложения", ош)
	}
	прилож.Пуск()
}