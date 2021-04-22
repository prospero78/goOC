// Package constexpres -- пакет предоставляет тип вычислимого выражения
//   на этапе компиляции для констант
package constexpres

import (
	"log"

	"github.com/prospero78/goOC/internal/types"
)

// TConstExpression -- операции с вычислимым выражением констант
type TConstExpression struct {
	listWord []types.IWord // Члены выражения
	word     types.IWord   // Фактическое значение
}

// New -- возвращает новый *TExpression
func New() *TConstExpression {
	return &TConstExpression{
		listWord: make([]types.IWord, 0),
	}
}

// GetWords -- возвращает слова выражения
func (sf TConstExpression) GetWords() []types.IWord {
	return sf.listWord
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
func (sf *TConstExpression) AddWord(word types.IWord) {
	if word == nil {
		log.Panicf("TConstExpression.AddWord(): word==nil\n")
	}
	sf.listWord = append(sf.listWord, word)
}

// GetWord -- возвращает хранимое слово значения
func (sf *TConstExpression) GetWord() types.IWord {
	return sf.word
}
