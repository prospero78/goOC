// Package constexpres -- пакет предоставляет тип вычислимого выражения
//   на этапе компиляции для констант
package constexpres

import (
	"log"
	"oc/internal/app/scanner/word"
)

// TConstExpression -- операции с вычислимым выражением констант
type TConstExpression struct {
	poolWord []*word.TWord // Члены выражения
	word     *word.TWord    // Фактическое значение
}

// New -- возвращает новый *TExpression
func New() *TConstExpression {
	return &TConstExpression{
		poolWord: make([]*word.TWord, 0),
	}
}

// GetWords -- возвращает слова выражения
func (sf TConstExpression) GetWords() []*word.TWord {
	return sf.poolWord
}

// SetType -- устанавливает значение типа константы
func (sf *TConstExpression) SetType(strType string) {
	if strType == "" {
		log.Panicf("TConstExpression.SetType(): strType==''\n")
	}
	if sf.word.GetType() != "" {
		log.Panicf("TConstExpression.SetType(): type(%v) already set, strType=%v\n", sf.word.GetType(), strType)
	}
	sf.word.SetType(strType)
}

// AddWord -- добавляет слово в выражение
func (sf *TConstExpression) AddWord(word *word.TWord) {
	if word == nil {
		log.Panicf("TConstExpression.AddWord(): word==nil\n")
	}
	sf.poolWord = append(sf.poolWord, word)
}

// GetWord -- возвращает хранимое слово значения
func (sf *TConstExpression)GetWord()*word.TWord{
	return sf.word
}
