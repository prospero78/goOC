package scanner

import (
	"oc/internal/app/scanner/stringsource"
	"oc/internal/log"
	"strings"
)

/*
	Пакет предоставляет сканер для разбора исходника
*/

// TScanner -- операции с исходником
type TScanner struct {
	log     *log.TLog
	poolStr []*stringsource.TStringSource
	poolWord []*word.TWord
}

// New -- возвращает новый *TScanner
func New() *TScanner {
	log := log.New("TScanner", log.DEBUG)
	log.Debugf("New")
	return &TScanner{
		log:     log,
		poolStr: make([]*stringsource.TStringSource, 0),
	}
}

// Scan -- сканирует исходник, разбивает на необходимые структуры
func (sf *TScanner) Scan(strSource string) {
	sf.log.Debugf("Scan")
	poolString := strings.Split(strSource, "\n")
	sf.log.Debugf("Run", "lines=", len(poolString))
	for num, str := range poolString {
		ss := stringsource.New(num+1, str)
		sf.poolStr = append(sf.poolStr, ss)
	}
	// Каждую строку разбить на слова
	for _, ss := range sf.poolStr {
		sf.scanString(ss)
	}
}

// Сканирует каждую строку, разбивает на слова
func (sf *TScanner) scanString(ss *stringsource.TStringSource) {
	str := ss.Val()
	isWord := false
	word:=""
	for _, rune := range str {
		lit := string(rune)
		switch {
		case lit==" ":
			isWord = false

		case strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", lit):
			isWord=true
		default:
			sf.log.Panicf("TScanner.scanString(): unknown lit(%v)\n", lit)
		}
	}
}
