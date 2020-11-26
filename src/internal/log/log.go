package log

/*
	Модуль предоставляет служебное логирование для всего приложения.
*/

import (
	"fmt"
	"time"
)

//TLog -- операции с логами типов
type TLog struct {
	pref string //Постоянный префикс для вывода
	mode int    //Режим работы типа
}

const (
	// DEBUG -- режим отладки приложения
	DEBUG = iota
	// INFO -- режим информирования приложения
	INFO
	// ERROR -- режим вывода только ошибок приложения
	ERROR
)

// New -- возвращает указатель на новый TLog
func New(pref string, mode int) (log *TLog) {
	log = &TLog{
		pref: pref,
		mode: mode,
	}
	return log
}

// Debugf -- печатает сообщение, если установлен режим отладки
func (sf *TLog) Debugf(method string, poolArg ...interface{}) {
	if sf.mode <= DEBUG {
		tineNow := time.Now().Format("2006-01-02 15:04:05.000")
		fmt.Printf("DEBU %v %v.%v():", tineNow, sf.pref, method)
		for _, arg := range poolArg {
			fmt.Printf(" %v", arg)
		}
		fmt.Printf("\n")
	}
}

// Infof -- печатает информацию, если установлен режим информации
func (sf *TLog) Infof(method string, poolArg ...interface{}) {
	if sf.mode <= INFO {
		tineNow := time.Now().Format("2006-01-02 15:04:05.000")
		fmt.Printf("INFO %v %v.%v():", tineNow, sf.pref, method)
		for _, arg := range poolArg {
			fmt.Printf(" %v", arg)
		}
		fmt.Printf("\n")
	}
}

// Panicf -- генерирует панику по требованию
func (sf *TLog) Panicf(method string, poolArg ...interface{}) {
	tineNow := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("PANIC %v %v.%v():", tineNow, sf.pref, method)
	for _, arg := range poolArg {
		fmt.Printf(" %v", arg)
	}
	fmt.Printf("\n")
	panic("")
}
