package app

/*
	Модуль предоставляет тип Оберон-компилятора.
	Здесь начинается вся работа.
*/

import (
	"io/ioutil"
	"oc/internal/app/scanner"
	"oc/internal/app/sectionset"
	"oc/internal/log"
	"os"
)

//TOc -- Оберон-компилятор (главный тип приложения)
type TOc struct {
	log     *log.TLog
	scanner *scanner.TScanner
	section *sectionset.TSectionSet
}

//New -- взвращает указатель на новый ТОк
func New(vers, build, data string) (oc *TOc, err error) {
	oc = &TOc{
		log:     log.New("TOc", log.DEBUG),
		scanner: scanner.New(),
		section: sectionset.New(),
	}
	oc.log.Debugf("New", "Создание типа компилятора")
	return oc, nil
}

//Run -- запуск компилтора после создания объекта компилятора
func (sf *TOc) Run() {
	sf.log.Debugf("Run")
	lenArgs := len(os.Args)
	if lenArgs < 2 {
		sf.log.Panicf("Run()", "for compile plis set file name\n")
	}
	fileName := os.Args[1]
	sf.log.Debugf("Run()", "fileName=", fileName)
	binSource, err := ioutil.ReadFile(fileName)
	if err != nil {
		sf.log.Panicf("Run()", "Error in read file", fileName, err)
	}
	strSource := string(binSource)
	sf.scanner.Scan(strSource)

	// Разбить по секциям
	sf.section.Split(sf.scanner)
}
