package main

/*
	Модуль определяет точку входа всего приложения
*/

import (
	"oc/internal/app"
	"oc/internal/log"
	mSs "oc/internal/splashscreen"
)

var (
	vers  = "0.0.8"
	build = "0095"
	data  = "2020-12-01 11:43:23"
	mode  = log.DEBUG
)

func main() {
	mode := log.DEBUG
	mSs.Print(vers, build, data)
	log := log.New("main", mode)
	log.Debugf("main", "Начало работы компилятора")
	app, err := app.New(vers, build, data)
	if err != nil {
		log.Panicf("main", "Запуск приложения", err)
	}
	app.Run()
}
