// Package word -- предоставляет тип слова. Содержит само слово и его атрибуты.
package word

import (
	"log"
	"oc/internal/app/sectionset/module/keywords"
	"strconv"
	"strings"
)

// TWord -- операции со словом
type TWord struct {
	pos      int    // Позиция в строке
	numStr   int    // Номер строки
	word     string // Само слово
	keywords *keywords.TKeywords
	strType  string // Строковое представление типа
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
		pos:      pos,
		numStr:   numStr,
		word:     val,
		keywords: keywords.Keys,
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

// Проверяет часть имени на строгое соответствие требованиям к имени
func (sf *TWord) isName(name string) bool {
	lit := string([]rune(name)[0])
	if res := sf.isLetter(lit); res != 0 { // Проверка на недопустимую первую литеру
		return false
	}
	for _, rune := range []rune(name) {
		lit = string(rune)
		if res := sf.isLetter(lit); !(res == 0 || res == 2) {
			return false
		}
	}
	return true
}

// IsCompoundName -- проверяет, что имя является составным
func (sf *TWord) IsCompoundName() bool {
	poolName := strings.Split(sf.word, ".")
	for _, name := range poolName {
		if !sf.isName(name) {
			return false
		}
	}
	return true
}

// NumStr -- возвращает номер строки
func (sf *TWord) NumStr() int {
	return sf.numStr
}

// IsInt -- проверяет, что слово является целым числом
func (sf *TWord) IsInt() bool {
	if string(sf.word[0]) == "_" {
		return false
	}
	sf.word = strings.ReplaceAll(sf.word, "_", "")
	_, err := strconv.Atoi(sf.word)
	return err == nil
}

// IsReal -- проверяет, что слово является вещественным числом
func (sf *TWord) IsReal() bool {
	_, err := strconv.ParseFloat(sf.word, 64)
	return err == nil
}

// IsString -- проверяет, что слово является строкой
func (sf *TWord) IsString() bool {
	run := []rune(sf.word)
	litBeg := string(run[0])
	if litBeg != "\"" {
		return false
	}
	litEnd := string(run[len(run)-1:])
	if litEnd != "\"" {
		return false
	}
	return true
}

// IsBool -- проверяет, что слово является булевым числом
func (sf *TWord) IsBool() bool {
	if sf.keywords.IsKey("TRUE", sf.word) {
		return true
	}
	if sf.keywords.IsKey("FALSE", sf.word) {
		return true
	}
	return false
}

// SetType -- устанавливает значение типа слова
func (sf *TWord) SetType(strType string) {
	if strType == "" {
		log.Panicf("TWord.SetType(): strType==''\n")
	}
	if sf.strType != "" {
		if sf.strType != strType {
			log.Panicf("TWord.SetType(): type(%v)!=strType(%v)\n", sf.strType, strType)
		}
		return
	}
	if sf.strType != "" {
		log.Panicf("TWord.SetType(): type(%v) already set, strType=%v\n", sf.strType, strType)
	}
	sf.strType = strType
}

// GetType -- возвращает хранимое значение типа
func (sf *TWord) GetType() string {
	return sf.strType
}
