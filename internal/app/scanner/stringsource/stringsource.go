package stringsource

import "log"

/*
	Пакет предоставляет тип для хранения строки исходника и её номера.
*/

// TStringSource -- операции с строкой исходника
type TStringSource struct {
	num int    // Номер строки
	val string // Значение строки
}

// New -- возвращает новый *TStringSource
func New(num int, val string) *TStringSource {
	if num < 1 {
		log.Panicf("New(): num(%v)<1\n", num)
	}
	ss := &TStringSource{
		num: num,
		val: val,
	}
	return ss
}

// Num -- возвращает номер строки
func (sf *TStringSource) Num() int {
	return sf.num
}

// Val -- возвращает хранимую строку
func (sf *TStringSource) Val() string {
	return sf.val
}
