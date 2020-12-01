package app

/*
	Модуль предоставляет тип Оберон-компилятора.
	Здесь начинается вся работа.
*/

import (
	"io/ioutil"
	"log"
	"oc/internal/app/modules"
	"oc/internal/app/scanner"
	"oc/internal/app/sectionset"
	"os"
	"strings"
)

//TOc -- Оберон-компилятор (главный тип приложения)
type TOc struct {
	scanner  *scanner.TScanner       // сканнер слов в модуле
	section  *sectionset.TSectionSet // разбивщик модуля на секции
	modules  *modules.TModules       // Набор модулей для компиляции
	filePath string                  // Имя главного файла для компиляции
}

//New -- взвращает указатель на новый ТОк
func New(vers, build, data string) (oc *TOc, err error) {
	lenArgs := len(os.Args)
	if lenArgs < 2 {
		log.Panicf("app.go/New(): for compile plis set file name\n")
	}
	filePath := os.Args[1]
	oc = &TOc{
		scanner:  scanner.New(),
		section:  sectionset.New(),
		filePath: filePath,
		modules:  modules.New(),
	}
	log.Printf("app.go/New(): создание типа компилятора")
	return oc, nil
}

//Run -- запуск компилтора после создания объекта компилятора
func (sf *TOc) Run() {
	log.Printf("TOc.Run(): fileName=%v\n", sf.filePath)
	binSource, err := ioutil.ReadFile(sf.filePath)
	if err != nil {
		log.Panicf("TOc.Run(): error in read file(%v)\n\t%v", sf.filePath, err)
	}
	strSource := string(binSource)
	sf.scanner.Scan(strSource)

	// Разбить по секциям
	sf.section.Split(sf.scanner)
	// Проверить, что имя главного модуля совпадает с имнем файла
	poolPath := strings.Split(sf.filePath, "/")
	fileName := poolPath[len(poolPath)-1]
	fileName = fileName[:len(fileName)-3]
	if fileName != sf.section.ModuleName() {
		log.Panicf("TOc.Run(): fileName(%v)!=mainName(%v)\n", fileName, sf.section.ModuleName())
	}
	sf.modules.SetMain(fileName, sf.section.Module())
}
