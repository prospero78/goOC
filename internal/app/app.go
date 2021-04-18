package app

/*
	Модуль предоставляет тип Оберон-компилятора.
	Здесь начинается вся работа.
*/

import (
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/prospero78/goOC/internal/app/modules"
	"github.com/prospero78/goOC/internal/app/scanner"
	"github.com/prospero78/goOC/internal/app/sectionset"
)

// TOc -- Оберон-компилятор (главный тип приложения)
type TOc struct {
	scanner  *scanner.TScanner       // сканнер слов в модуле
	section  *sectionset.TSectionSet // разбивщик модуля на секции
	modules  *modules.TModules       // Набор модулей для компиляции
	filePath string                  // Имя главного файла для компиляции
	path     string                  // общий путь к файлу
}

// New -- взвращает указатель на новый ТОк
func New(vers, build, data, filePath string) (oc *TOc, err error) {
	logrus.Debugln("app.go/New(): создание типа компилятора")
	oc = &TOc{
		scanner:  scanner.New(),
		section:  sectionset.New(),
		filePath: filePath,
		modules:  modules.New(),
	}
	return oc, nil
}

// Run -- запуск компилтора после создания объекта компилятора
func (sf *TOc) Run() {
	logrus.WithField("filePath", sf.filePath).Debugln("TOc.Run()")
	strSource := sf.readFile(sf.filePath)
	poolName := strings.Split(sf.filePath, "/")
	nameMod := poolName[len(poolName)-1]
	nameMod = nameMod[:len(nameMod)-3]
	sf.scanner.Scan(nameMod, strSource)
	// Разбить по секциям
	sf.section.Split(sf.scanner)
	nameMain := sf.section.Module().Name()
	sf.path = sf.filePath[:len(sf.filePath)-len(nameMain+".o7")]
	sf.checkModuleName(sf.filePath, nameMain)
	sf.modules.SetMain(nameMain)
	sf.modules.AddModule(sf.section.Module())
	sf.getImport(nameMain)
	logrus.WithField("all modules", sf.modules.Len()).Debugln("Toc.Run()")
	sf.modules.ProcessConstant()
}

// По требованию читает файл, возвращает содержимое
func (sf *TOc) readFile(filePath string) (strSource string) {
	binSource, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.WithError(err).WithField("fileName", sf.filePath).Panicf("TOc.readFile(): error in read file\n")
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
		logrus.Panicf("TOc.checkModuleName(): fileName(%v)!=moduleName(%v)\n", fileName, moduleName)
	}
}

func (sf *TOc) getImport(nameModule string) {
	// Получить остальные импорты модулей.
	modules := sf.section.GetImport()
	for _, module := range modules {
		if module == nil {
			logrus.Panicf("TOc.Run(): module==nil\n")
		}
		sf.scanModule(module.Name())
	}
}

// Готовит параметры под новый сканер
func (sf *TOc) scanModule(moduleName string) {
	if sf.modules.IsExist(moduleName) {
		return
	}
	filePath := sf.path + moduleName + ".o7"
	sf.scanner = scanner.New()
	sf.section = sectionset.New()
	strSource := sf.readFile(filePath)
	sf.scanner.Scan(moduleName, strSource)
	sf.section.Split(sf.scanner)
	module := sf.section.Module()
	sf.checkModuleName(filePath, module.Name())
	sf.modules.AddModule(module)
}
