// Package main -- определяет точку входа всего приложения
package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/prospero78/goOC/cmd/gooc/argoc"
	"github.com/prospero78/goOC/internal/app"
	mSs "github.com/prospero78/goOC/internal/splashscreen"
)

var (
	vers  = "0.0.8"
	build = "0100"
	data  = "2021-04-21 22:34:25"
)

func main() {
	argOc := argoc.GetArgOc()
	mSs.Print(vers, build, data)
	log.Debugf("main(): Начало работы компилятора")
	app, err := app.New(vers, build, data, argOc.FileName())
	if err != nil {
		log.WithError(err).Panicf("main(): Запуск приложения")
	}
	if err := app.Run(); err != nil {
		log.WithError(err).Panicf("main(): in run app\n")
	}
}
