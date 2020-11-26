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
		Vers:  "0.0.6",
		Build: "0070",
		Data:  "2020-11-26 11:56:00",
		Mode:  mode,
	}
	mSs.Print(param)
	log := log.New("main", mode)
	log.Debugf("main", "Начало работы компилятора")
	app, err := app.New(param)
	if err != nil {
		log.Panicf("main", "Запуск приложения", err)
	}
	app.Пуск()
}
