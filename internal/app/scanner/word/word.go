package word

import (
	"log"
	"strings"
)

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

// Проверяет, что есть конкретная литера
func (sf *TWord) isLetter(lit string) (res int) {
	if len(sf.word) == 0 {
		log.Panicf("TWord.isLetter(): word==''")
	}
	if strings.Contains("_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", lit) { // en_En.UTF-8
		return 0
	}
	if strings.Contains("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя", lit) { // ru_RU.UTF-8
		return 0
	}
	if strings.Contains(".@!", lit) { // Если допустимые литеры
		return 1
	}
	if strings.Contains("1234567890", lit) { // Если цифры
		return 2
	}
	return 3
}

// IsName -- проверяет слово на строгое соответствие требованиям к имени
func (sf *TWord) IsName() bool {
	lit := string([]rune(sf.word)[0])
	if res := sf.isLetter(lit); res != 0 { // Проверка на недопустимую первую литеру
		return false
	}
	for _, rune := range []rune(sf.word) {
		lit = string(rune)
		if res := sf.isLetter(lit); !(res == 0 || res == 2) {
			return false
		}
	}
	return true
}

// NumStr -- возвращает номер строки
func (sf *TWord)NumStr()int{
	return sf.numStr
}
