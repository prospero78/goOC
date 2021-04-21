// Package word -- предоставляет тип слова. Содержит само слово и его атрибуты.
package word

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/prospero78/goOC/internal/app/scanner/word/litpos"
	"github.com/prospero78/goOC/internal/app/scanner/word/numstr"
	"github.com/prospero78/goOC/internal/app/scanner/word/strword"
	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/types"
)

const (
	litEng   = "_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	litRu    = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	litDigit = "1234567890"
)

// TWord -- операции со словом
type TWord struct {
	pos      types.IPosFix   // Позиция в строке
	numStr   types.INumStr   // Номер строки
	word     types.IStrWord  // Само слово
	keywords types.IKeywords // Ссылка на глобальный объект ключевых слов
	strType  string          // Строковое представление типа
	module   *types.AModule  // Имя модуля для слова
}

// New -- возвращает новый *TWord
func New(numStr types.ANumStr, pos types.APos, strWord types.AWord) (*TWord, error) {
	_pos, err := litpos.New(pos)
	if err != nil {
		return nil, fmt.Errorf("word.go/New(): in create IPosFix, err=%w", err)
	}
	_numStr, err := numstr.New(numStr)
	if err != nil {
		return nil, fmt.Errorf("word.go/New(): in create INumStr, err=%w", err)
	}
	_word, err := strword.New(strWord)
	if err != nil {
		return nil, fmt.Errorf("word.go/New(): in create IStrWord, err=%w", err)
	}
	word := &TWord{
		pos:      _pos,
		numStr:   _numStr,
		word:     _word,
		keywords: keywords.GetKeys(),
	}
	return word, nil
}

// Word -- возвращает хранимое слово
func (sf *TWord) Word() types.AWord {
	return sf.word.Get()
}

// Проверяет, что есть конкретная литера
func (sf *TWord) isLetter(lit string) (res int) {
	if sf.word.Len() == 0 {
		logrus.Panicf("TWord.isLetter(): word==''")
	}
	switch {
	case strings.Contains(litEng, lit): // en_En.UTF-8
		return 0
	case strings.Contains(litRu, lit): // ru_RU.UTF-8
		return 0
	case strings.Contains(".@!", lit): // Если допустимые литеры
		return 1
	case strings.Contains(litDigit, lit): // Если цифры
		return 2
	}
	return 3
}

// IsName -- проверяет слово на строгое соответствие требованиям к имени
func (sf *TWord) IsName() bool {
	firstLit := string([]rune(sf.word.Get())[0])
	if res := sf.isLetter(firstLit); res != 0 { // Проверка на недопустимую первую литеру
		return false
	}
	for _, rune := range sf.word.Get() {
		nextLit := string(rune)
		if res := sf.isLetter(nextLit); !(res == 0 || res == 2) {
			return false
		}
	}
	return true
}

// Проверяет часть имени на строгое соответствие требованиям к имени
func (sf *TWord) checkPartName(name string) bool {
	firstLit := string([]rune(name)[0])
	if res := sf.isLetter(firstLit); res != 0 { // Проверка на недопустимую первую литеру
		return false
	}
	for _, rune := range name {
		nextLit := string(rune)
		if res := sf.isLetter(nextLit); !(res == 0 || res == 2) {
			return false
		}
	}
	return true
}

// IsCompoundName -- проверяет, что имя является составным
func (sf *TWord) IsCompoundName() bool {
	poolName := strings.Split(string(sf.word.Get()), ".")
	for _, name := range poolName {
		if !sf.checkPartName(name) {
			return false
		}
	}
	return true
}

// NumStr -- возвращает номер строки
func (sf *TWord) NumStr() types.ANumStr {
	return sf.numStr.Get()
}

// IsInt -- проверяет, что слово является целым числом
func (sf *TWord) IsInt() bool {
	if string(sf.word.Get()[0]) == "_" {
		return false
	}
	str := strings.ReplaceAll(string(sf.word.Get()), "_", "")
	_, err := strconv.Atoi(str)
	return err == nil
}

// IsReal -- проверяет, что слово является вещественным числом
func (sf *TWord) IsReal() bool {
	_, err := strconv.ParseFloat(string(sf.word.Get()), 64)
	return err == nil
}

// IsString -- проверяет, что слово является строкой (должно быть в кавычках)
func (sf *TWord) IsString() bool {
	run := []rune(sf.word.Get())
	litFirst := string(run[0])
	if litFirst != "\"" {
		return false
	}
	litEnd := string(run[len(run)-1:])
	return litEnd == "\""
}

// IsBool -- проверяет, что слово является булевым числом
func (sf *TWord) IsBool() bool {
	switch {
	case sf.keywords.IsKey("TRUE", sf.word.Get()):
		return true
	case sf.keywords.IsKey("FALSE", sf.word.Get()):
		return true
	}
	return false
}

// SetType -- устанавливает значение типа слова
func (sf *TWord) SetType(strType string) {
	switch {
	case strType == "":
		logrus.Panicf("TWord.SetType(): strType==''\n")
	case sf.strType != "":
		if sf.strType == strType {
			return
		}
		logrus.Panicf("TWord.SetType(): type(%v)!=strType(%v)\n", sf.strType, strType)
	}
	sf.strType = strType
}

// GetType -- возвращает хранимое значение типа
func (sf *TWord) GetType() string {
	return sf.strType
}

// SetModule -- устанавливает имя модуля
func (sf *TWord) SetModule(module *types.AModule) {
	if *module == "" {
		logrus.Panicf("TWord.SetModule(): module=''\n")
	}
	sf.module = module
}

// Module -- возвращает хранимое имя модуля
func (sf *TWord) Module() *types.AModule {
	return sf.module
}

// Pos -- возвращает позицию в строке
func (sf *TWord) Pos() types.APos {
	return sf.pos.Get()
}
