package word

import "log"

/*
	Пакет предоставляет тип слова.
	Содержит само слово и его атрибуты.
*/

// TWord -- операции со словом
type TWord struct {
	pos    int    // Позиция в строке
	numStr int    // Номер строки
	word   string // Само слово
}

// New -- возвращает новый *TWord
func New(numStr, pos int, val string) *TWord {
	{ // Предусловия
		if numStr < 1 {
			log.Panicf("word.go/New(): numStr(%v)<1\n", numStr)
		}
		if pos < 0 {
			log.Panicf("word.go/New(): pos(%v)<0\n", pos)
		}
		if val == "" {
			log.Panicf("word.go/New(): val==''\n")
		}
	}
	word := &TWord{
		pos:    pos,
		numStr: numStr,
		word:   val,
	}
	return word
}

// Word -- возвращает хранимое слово
func (sf *TWord) Word() string {
	return sf.word
}
