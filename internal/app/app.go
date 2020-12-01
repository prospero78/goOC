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
	//"oc/internal/app/sectionset/module"
	"os"
)

//TOc -- Оберон-компилятор (главный тип приложения)
type TOc struct {
	scanner  *scanner.TScanner       // сканнер слов в модуле
	section  *sectionset.TSectionSet // разбивщик модуля на секции
	modules  *modules.TModules       // Набор модулей для компиляции
	filePath string                  // Имя главного файла для компиляции
	path     string                  // общий путь к файлу
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
	strSource := sf.readFile(sf.filePath)
	sf.scanner.Scan(strSource)
	// Разбить по секциям
	sf.section.Split(sf.scanner)
	nameMain := sf.section.Module().Name()
	sf.path = sf.filePath[:len(sf.filePath)-len(nameMain+".o7")]
	sf.checkModuleName(sf.filePath, nameMain)
	sf.modules.SetMain(nameMain, sf.section.Module())
	sf.getImport(nameMain)
	log.Printf("Toc.Run(): all modules=%v\n", sf.modules.Len())
}

// По требованию читает файл, возвращает содержимое
func (sf *TOc) readFile(filePath string) (strSource string) {
	binSource, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("TOc.readFile(): error in read file(%v)\n\t%v", filePath, err)
	}
	strSource = string(binSource)
	return strSource
}

// Проверяет имя модуля на соответствие имени файла
func (sf *TOc) checkModuleName(fileName, moduleName string) {
	// Проверить, что имя главного модуля совпадает с имнем файла
	fileName = fileName[len(sf.path):]
	fileName = fileName[:len(fileName)-3]
	if fileName != moduleName {
		log.Panicf("TOc.checkModuleName(): fileName(%v)!=moduleName(%v)\n", fileName, moduleName)
	}
}

func (sf *TOc) getImport(nameModule string) {
	// Получить остальные импорты модулей.
	modules := sf.section.GetImport()
	for _, module := range modules {
		if module == nil {
			log.Panicf("TOc.Run(): module==nil\n")
		}
		sf.scanModule(module.Name())
	}
}

// Готовит параметры под новый сканер
func (sf *TOc) scanModule(moduleName string) {
	filePath := sf.path + moduleName + ".o7"
	sf.scanner = scanner.New()
	sf.section = sectionset.New()
	strSource := sf.readFile(filePath)
	sf.scanner.Scan(strSource)
	sf.section.Split(sf.scanner)
	module := sf.section.Module()
	sf.checkModuleName(filePath, module.Name())
	sf.modules.AddModule(module)
}
