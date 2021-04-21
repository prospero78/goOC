// Package argoc -- параметры запуска компилятора
package argoc

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/prospero78/goOC/internal/types"
)

// TArgOc -- операции с параметрами запуска
type TArgOc struct {
	fileName string
}

var (
	argOc *TArgOc
)

// GetArgOc -- возвращает объект параметров запуска
func GetArgOc() types.IArgOc {
	if argOc != nil {
		return argOc
	}
	argOc = &TArgOc{}
	argOc.init()
	return argOc
}

// Инициализирует параметры компилятора
func (sf *TArgOc) init() {
	debug := os.Getenv("DEBUG")
	if debug != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	args := os.Args
	if len(args) != 2 {
		logrus.Panicf(`TArgOc.init(): for run compiler
	gooc <file_name>.o7
	
Example:
	gooc Main.o7`)
	}
	sf.fileName = args[1]
}

// FileName -- возвращает имя файла для компиляции
func (sf *TArgOc) FileName() string {
	return sf.fileName
}
