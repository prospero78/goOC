// Package main -- определяет точку входа всего приложения
package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/prospero78/goOC/internal/app"
	mSs "github.com/prospero78/goOC/internal/splashscreen"
)

var (
	vers  = "0.0.8"
	build = "0095"
	data  = "2020-12-01 11:43:23"
)

func main() {
	mSs.Print(vers, build, data)
	log.Debugf("main(): hачало работы компилятора")
	app, err := app.New(vers, build, data)
	if err != nil {
		log.WithError(err).Panicf("main(): Запуск приложения")
	}
	app.Run()
}
