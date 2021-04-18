// Package calcword -- рассчитывает тип и значение слова (если возможно)
package calcword

import (
	"log"
	"github.com/prospero78/goOC/internal/app/scanner/word"
	"strings"
)

// TCalcWord -- операции по расчёту типа и значения слова
type TCalcWord struct {
	word *word.TWord
}

// New -- возвращает новый *TCalcWord
func New() *TCalcWord {
	return &TCalcWord{}
}

// RecognizeType -- распознаёт тип слова
func (sf *TCalcWord) RecognizeType(word *word.TWord) (name, operation string) {
	switch {
	case word.IsInt(): // Если целое
		word.SetType("INTEGER")
	case word.IsReal(): // Если вещественное
		word.SetType("REAL")
	case word.IsString(): // Если строка
		word.SetType("ARRAY OF CHAR")
	case word.IsBool(): // Если булево
		word.SetType("BOOLEAN")
	case word.Word() == "(": // Если начало выражения
		word.SetType("(")
	case word.Word() == ")": // Если конец выражения
		word.SetType(")")
	case word.IsName(): // Если присвоение из другого слова
		return word.Word(), "type, val" // Запросить тип и значение этого слова
	case word.IsCompoundName(): // Имя состоит из нескольких частей
		poolName := strings.Split(word.Word(), ".")
		// Проверить, что "Модуль:имя"
		if len(poolName) == 2 {
			return word.Word(), "module, name, type, val"
		}
	case word.Word() == "+": // Операция "+"
		word.SetType("+")
	case word.Word() == "-": // Операция "-"
		word.SetType("-")
	case word.Word() == "/": // Операция РАЗДЕЛИТЬ
		word.SetType("/")
	case word.Word() == "*": // Операция "*"
		word.SetType("*")
	default:
		log.Panicf("TModules.checkTypeConstant(): unknown type for constante %v\n", word.Word())
	}
	return "", ""
}
