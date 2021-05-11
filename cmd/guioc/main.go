// Package main -- графическая оболочка для компилятора
package main

import (
	"github.com/prospero78/goOC/internal/gui"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infoln("main.go/main(): старт компилятора")
	oc := gui.New()
	oc.Run()
}
