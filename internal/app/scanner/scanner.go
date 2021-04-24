package scanner

/*
	Пакет предоставляет сканер для разбора исходника
*/

import (
	// "fmt"
	"fmt"
	"strings"

	"github.com/prospero78/goOC/internal/app/scanner/isletter"
	"github.com/prospero78/goOC/internal/app/scanner/stringsource"
	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/types"
	"github.com/sirupsen/logrus"
)

const (
	litEng   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	litRu    = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	litDigit = "1234567890"
)

// TScanner -- операции с исходником
type TScanner struct {
	listStr  []*stringsource.TStringSource
	listWord []types.IWord
	listRune []rune        // Текущая строка исходника
	pos      types.APos    // Позиция руны в строке
	num      types.ANumStr // Номер строки
}

// New -- возвращает новый *TScanner
func New() *TScanner {
	return &TScanner{
		listStr:  make([]*stringsource.TStringSource, 0),
		listWord: make([]types.IWord, 0),
		num:      1,
	}
}

// Scan -- сканирует исходник, разбивает на необходимые структуры
func (sf *TScanner) Scan(nameMod types.AModule, strSource string) {
	// log.Printf("Scan")
	listString := strings.Split(strSource, "\n")
	for num, str := range listString {
		ss := stringsource.New(num+1, str)
		sf.listStr = append(sf.listStr, ss)
	}
	sf.scanString(strSource)

	// Присовить всем словам имя модуля
	for _, word := range sf.listWord {
		if err := word.SetModule(&nameMod); err != nil {
			logrus.WithError(err).Panicf("TScanner.Scan(): in set module for word\n")
		}
	}
	// log.Printf("TScanner.Run(): lines=%v word=%v\n", len(poolString), len(sf.poolWord))
	// for _, word := range sf.poolWord {
	// 	fmt.Printf("%v\t", word.Word())
	// }
}

// Сканирует каждую строку, разбивает на слова
func (sf *TScanner) scanString(strSource string) (err error) {
	sf.listRune = []rune(strSource)
	for len(sf.listRune) > 0 {
		lit := string(sf.listRune[0])
		if isletter.Check(lit) { // Если начало слова
			sf.getWord()
			continue
		}
		if sf.isCommentOpen() { // Проверка на комментарий
			continue
		}
		if sf.isCommentClose() { // Проверка на комментарий
			continue
		}
		if sf.isTerminalEmpty() { // Разделитель слов
			continue
		}
		if sf.isTerminalExp() { // Разделитель выражений
			continue
		}
		if sf.isDeclaration() { // Проверка на объявление/присовение
			continue
		}
		if sf.isCompare() { // Проверка на сравнение
			continue
		}
		if sf.isString() { // Проверка на строку
			continue
		}
		isBracket, err := sf.isBracket()
		if err != nil {
			return fmt.Errorf("TScanner.scanString(): при проверке круглых скобок\n\t%w", err)
		}
		if isBracket { // Проверка на скобки
			continue
		}
		return fmt.Errorf("TScanner.scanString(): неизвестная литера (%q)\n", string(sf.listRune[0]))
	}
	return nil
}

// Проверяет, что литера является скобкой
func (sf *TScanner) isBracket() (isBracket bool, err error) {
	isExp := false
	lit := string(sf.listRune[0])
	if strings.Contains("()", lit) {
		sf.pos++
		if err = sf.addWord(types.AWord(lit)); err != nil {
			return false, fmt.Errorf("TScabber.isBracket(): in add round bracket\n\t%w", err)
		}
		isExp = true
	}
	if isExp {
		sf.listRune = sf.listRune[1:]
		return true, nil
	}
	return false, nil
}

// Выбирает закрытие комментария
func (sf *TScanner) isCommentClose() bool {
	lit := string(sf.listRune[0])
	if lit != "*" {
		return false
	}
	sf.pos++
	sf.listRune = sf.listRune[1:]
	lit = string(sf.listRune[0])
	switch lit {
	case ")":
		sf.pos++
		sf.addWord("*)")
		sf.listRune = sf.listRune[1:]
		return true
	default:
		sf.addWord("*")
		return true
	}
}

// Выбирает открытие комментария
func (sf *TScanner) isCommentOpen() bool {
	lit := string(sf.listRune[0])
	if lit != "(" {
		return false
	}
	sf.pos++
	sf.listRune = sf.listRune[1:]
	lit = string(sf.listRune[0])
	switch lit {
	case "*":
		sf.pos++
		sf.listRune = sf.listRune[1:]
		sf.addWord("(*")
		return true
	default:
		sf.addWord("(")
		return true
	}
}

// Проверяет, что литера является сравнением
func (sf *TScanner) isCompare() bool {
	isExp := false
	lit := string(sf.listRune[0])
	if strings.Contains("=", lit) {
		sf.pos++
		sf.addWord(types.AWord(lit))
		isExp = true
	}
	if isExp {
		sf.listRune = sf.listRune[1:]
		return true
	}
	return false
}

// Проверяет, что литера является объявлением/присвоением
func (sf *TScanner) isDeclaration() bool {
	lit := string(sf.listRune[0])
	if lit != ":" {
		return false
	}

	sf.pos++
	sf.listRune = sf.listRune[1:]
	lit = string(sf.listRune[0])
	if lit != "=" {
		sf.addWord(":")
		return true
	}

	sf.pos++
	sf.listRune = sf.listRune[1:]
	sf.addWord(":=")
	return true
}

// Проверяет, что литера является терминалом выражений
func (sf *TScanner) isTerminalExp() bool {
	isExp := false
	lit := string(sf.listRune[0])
	if strings.Contains(";*,-+/", lit) {
		sf.pos++
		sf.addWord(types.AWord(lit))
		isExp = true
	}
	if isExp {
		sf.listRune = sf.listRune[1:]
		return true
	}
	return false
}

// Проверяет, что литера является пустым терминалом слов
func (sf *TScanner) isTerminalEmpty() (res bool) {
	lit := string(sf.listRune[0])
	if strings.Contains(" \t", lit) {
		res = true
	}
	if lit == "\n" {
		sf.num++
		sf.pos = -1
		res = true
	}
	if res {
		sf.listRune = sf.listRune[1:]
		sf.pos++
		return true
	}
	return false
}

// Проверяет, что литера -- буква
func (sf *TScanner) isLetter() bool {
	if len(sf.listRune) == 0 {
		return false
	}
	lit := string(sf.listRune[0])
	if strings.Contains(litEng, lit) { // en_En.UTF-8
		return true
	}
	if strings.Contains(litRu, lit) { // ru_RU.UTF-8
		return true
	}
	if strings.Contains("_.@!", lit) { // Если допустимые литеры
		return true
	}
	if strings.Contains(litDigit, lit) { // Если цифры
		return true
	}
	return false
}

// Выбирает строку в кавычках
func (sf *TScanner) isString() bool {

	lit := string(sf.listRune[0])
	if lit != "\"" {
		return false
	}
	sf.pos++
	word := types.AWord("\"")
	sf.listRune = sf.listRune[1:]
	for len(sf.listRune) != 0 {
		lit := string(sf.listRune[0])
		if lit == "\"" {
			word += "\""
			sf.pos++
			sf.listRune = sf.listRune[1:]
			sf.addWord(word)
			return true
		}
		sf.pos++
		word += types.AWord(lit)
		sf.listRune = sf.listRune[1:]
	}
	return false
}

// Выбирает слово целиком из строки до разделителя
func (sf *TScanner) getWord() {
	word := types.AWord("")
	for sf.isLetter() {
		lit := string(sf.listRune[0])
		sf.pos++
		word += types.AWord(lit)
		sf.listRune = sf.listRune[1:]
	}
	sf.addWord(word)
}

// Добавляет слово в пул слов
func (sf *TScanner) addWord(wrd types.AWord) error {
	word, err := word.New(sf.num, sf.pos, wrd)
	if err != nil {
		return fmt.Errorf("TScanner.addWord(): in create word\n\t%w", err)
	}
	sf.listWord = append(sf.listWord, word)
	return err
}

// ListWord -- возвращает список слов после обработки
func (sf *TScanner) ListWord() (res []types.IWord) {
	res = make([]types.IWord, 0)
	res = append(res, sf.listWord...)
	return res
}
