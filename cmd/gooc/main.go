package main

/*
	Модуль определяет точку входа всего приложения
*/

import (
	"log"
	"oc/internal/app"
	"oc/internal/app/param"
	"oc/internal/log"
	mSs "oc/internal/splashscreen"
)

func main() {
	mode := log.DEBUG
	param := &param.ТПарам{
		Vers: "0.0.6",
		Build: "0070",
		Data:   "2020-08-31 22:03:00",
		Mode:  mode,
	}
	mSs.Print(param)
	log := log.Нов("main", mode)
	log.Debugf("main", "Начало работы компилятора")
	app, err := app.Нов(param)
	if err != nil {
		log.Panicf("main", "Запуск приложения", err)
	}
	app.Пуск()
}
