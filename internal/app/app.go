package app

/*
	Модуль предоставляет тип Оберон-компилятора.
	Здесь начинается вся работа.
*/

import (
	"oc/internal/app/module"
	"oc/internal/app/param"
	"oc/internal/log"
)

//TOc -- Оберон-компилятор (главный тип приложения)
type TOc struct {
	log  *log.TLog
	mode int //Режим работы компилятора
}

//New -- взвращает указатель на новый ТОк
func New(param *param.TParam) (ok *TOc, err error) {
	ok = &TOc{
		log:  log.New("TOc", param.Mode),
		mode: param.Mode,
	}
	ok.log.Debugf("New", "Создание типа компилятора")
	return ok, nil
}

//Run -- запуск компилтора после создания объекта компилятора
func (sf *TOc) Run() {
	sf.log.Debugf("Run")
	fileName := "Hello.o7"
	module, ош := module.Нов(fileName, sf.mode)
	if ош != nil {
		sf.log.Panicf("Run", "обработка модуля", ош)
	}
	_ = module.КодПолуч()
}
