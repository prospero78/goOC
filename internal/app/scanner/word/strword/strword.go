// Package strword -- неизменяемое строковое представление слова
package strword

import (
	"fmt"

	"github.com/prospero78/goOC/internal/types"
)

// TStrWord -- операции с неизменяемым строковым представлением слова
type TStrWord struct {
	val types.AWord
}

// New -- возвращает новый IWord
func New(word types.AWord) (sw types.IStrWord, err error) {
	if word == "" {
		return nil, fmt.Errorf("strword.go/New(): word is empty")
	}
	sw = &TStrWord{
		val: word,
	}
	return sw, nil
}

// Get -- возвращает хранимое значение
func (sf *TStrWord) Get() types.AWord {
	return sf.val
}

// Len -- возвращает длину слова
func (sf *TStrWord) Len() int {
	return len(sf.val)
}
