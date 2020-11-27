package scanner

import (
	//"fmt"
	"log"
	"oc/internal/app/scanner/stringsource"
	"oc/internal/app/scanner/word"
	"strings"
)

/*
	Пакет предоставляет сканер для разбора исходника
*/

// TScanner -- операции с исходником
type TScanner struct {
	poolStr   []*stringsource.TStringSource
	poolWord  []*word.TWord
	runSource []rune // Текущая строка исходника
	pos       int    //Позиция руны в строке
	num       int    // Номер строки
}

// New -- возвращает новый *TScanner
func New() *TScanner {
	return &TScanner{
		poolStr:  make([]*stringsource.TStringSource, 0),
		poolWord: make([]*word.TWord, 0),
		num:      1,
	}
}

// Scan -- сканирует исходник, разбивает на необходимые структуры
func (sf *TScanner) Scan(strSource string) {
	log.Printf("Scan")
	poolString := strings.Split(strSource, "\n")
	for num, str := range poolString {
		ss := stringsource.New(num+1, str)
		sf.poolStr = append(sf.poolStr, ss)
	}
	sf.scanString(strSource)

	log.Printf("TScanner.Run(): lines=%v word=%v\n", len(poolString), len(sf.poolWord))
	// for _, word := range sf.poolWord {
	// 	fmt.Printf("%v\t", word.Word())
	// }
}

// Сканирует каждую строку, разбивает на слова
func (sf *TScanner) scanString(strSource string) {
	sf.runSource = []rune(strSource)
	for len(sf.runSource) > 0 {
		if sf.isLetter() { // Если начало слова
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
		if sf.isBracket() { // Проверка на скобки
			continue
		}
		log.Panicf("TScanner.scanString(): неизвестная литера (%q)\n", string(sf.runSource[0]))
	}
}

// Проверяет, что литера является скобкой
func (sf *TScanner) isBracket() bool {
	isExp := false
	lit := string(sf.runSource[0])
	if strings.Contains("()", lit) {
		sf.pos++
		sf.addWord(lit)
		isExp = true
	}
	if isExp {
		sf.runSource = sf.runSource[1:]
		return true
	}
	return false
}

// Выбирает закрытие комментария
func (sf *TScanner) isCommentClose() bool {
	lit := string(sf.runSource[0])
	if lit != "*" {
		return false
	}
	sf.pos++
	sf.runSource = sf.runSource[1:]
	lit = string(sf.runSource[0])
	if lit != ")" {
		sf.addWord("*")
		return true
	}
	sf.pos++
	sf.addWord("*)")
	sf.runSource = sf.runSource[1:]
	return true
}

// Выбирает открытие комментария
func (sf *TScanner) isCommentOpen() bool {
	lit := string(sf.runSource[0])
	if lit != "(" {
		return false
	}
	sf.pos++
	sf.runSource = sf.runSource[1:]
	lit = string(sf.runSource[0])
	if lit != "*" {
		sf.addWord("(")
		return true
	}
	sf.pos++
	sf.runSource = sf.runSource[1:]
	sf.addWord("(*")
	return true
}

// Проверяет, что литера является сравнением
func (sf *TScanner) isCompare() bool {
	isExp := false
	lit := string(sf.runSource[0])
	if strings.Contains("=", lit) {
		sf.pos++
		sf.addWord(lit)
		isExp = true
	}
	if isExp {
		sf.runSource = sf.runSource[1:]
		return true
	}
	return false
}

// Проверяет, что литера является объявлением/присвоением
func (sf *TScanner) isDeclaration() bool {
	lit := string(sf.runSource[0])
	if lit != ":" {
		return false
	}

	sf.pos++
	sf.runSource = sf.runSource[1:]
	lit = string(sf.runSource[0])
	if lit != "=" {
		sf.addWord(":")
		return true
	}

	sf.pos++
	sf.runSource = sf.runSource[1:]
	sf.addWord(":=")
	return true
}

// Проверяет, что литера является терминалом выражений
func (sf *TScanner) isTerminalExp() bool {
	isExp := false
	lit := string(sf.runSource[0])
	if strings.Contains(";*,-+/", lit) {
		sf.pos++
		sf.addWord(lit)
		isExp = true
	}
	if isExp {
		sf.runSource = sf.runSource[1:]
		return true
	}
	return false
}

// Проверяет, что литера является пустым терминалом слов
func (sf *TScanner) isTerminalEmpty() (res bool) {
	lit := string(sf.runSource[0])
	if strings.Contains(" \t", lit) {
		res = true
	}
	if lit == "\n" {
		sf.num++
		sf.pos = -1
		res = true
	}
	if res {
		sf.runSource = sf.runSource[1:]
		sf.pos++
		return true
	}
	return false
}

// Проверяет, что литера -- буква
func (sf *TScanner) isLetter() bool {
	if len(sf.runSource) == 0 {
		return false
	}
	lit := string(sf.runSource[0])
	if strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", lit) { // en_En.UTF-8
		return true
	}
	if strings.Contains("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя", lit) { // ru_RU.UTF-8
		return true
	}
	if strings.Contains("_.@!", lit) { // Если допустимые литеры
		return true
	}
	if strings.Contains("1234567890", lit) { // Если цифры
		return true
	}
	return false
}

// Выбирает строку в кавычках
func (sf *TScanner) isString() bool {

	lit := string(sf.runSource[0])
	if lit != "\"" {
		return false
	}
	sf.pos++
	word := "\""
	sf.runSource = sf.runSource[1:]
	for len(sf.runSource) != 0 {
		lit := string(sf.runSource[0])
		if lit == "\"" {
			sf.pos++
			word += lit
			sf.runSource = sf.runSource[1:]
			sf.addWord(word)
			return true
		}
		sf.pos++
		word += lit
		sf.runSource = sf.runSource[1:]
	}
	return false
}

// Выбирает слово целиком из строки до разделителя
func (sf *TScanner) getWord() {
	word := ""
	for sf.isLetter() {
		lit := string(sf.runSource[0])
		sf.pos++
		word += lit
		sf.runSource = sf.runSource[1:]
	}
	sf.addWord(word)
}

// Добавляет слово в пул слов
func (sf *TScanner) addWord(wrd string) {
	word := word.New(sf.num, sf.pos, wrd)
	sf.poolWord = append(sf.poolWord, word)
}

// PoolWord -- возвращает пул слов после обработки
func (sf *TScanner) PoolWord() (res []*word.TWord) {
	res = make([]*word.TWord, 0)
	res = append(res, sf.poolWord...)
	return res
}
