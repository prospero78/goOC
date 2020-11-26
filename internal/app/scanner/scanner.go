package scanner

import (
	"oc/internal/log"
)

/*
	Пакет предоставляет сканер для разбора исходника
*/

// TScanner -- операции с исходником
type TScanner struct {
	log *log.TLog
}

// New -- возвращает новый *TScanner
func New() *TScanner {
	log := log.New("TScanner", log.DEBUG)
	log.Debugf("New")
	return &TScanner{
		log: log,
	}
}

// Scan -- сканирует исходник, разбивает на необходимые структуры
func (sf *TScanner) Scan(strSource string) {
	sf.log.Debugf("Scan")
	countLines := 0
	for _, lit := range strSource { // Запустить цикл просмотра строк
		switch string(lit) {
		case "\n", "\r": // Новая строка
			countLines++
		}
	}
	sf.log.Debugf("Run", "countLines=", countLines)
}
