package app

/*
	Модуль предоставляет тип Оберон-компилятора.
	Здесь начинается вся работа.
*/

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/prospero78/goOC/internal/app/modules"
	"github.com/prospero78/goOC/internal/app/scanner"
	"github.com/prospero78/goOC/internal/app/sectionset"
	"github.com/prospero78/goOC/internal/types"
)

// TOc -- Оберон-компилятор (главный тип приложения)
type TOc struct {
	scanner  *scanner.TScanner       // сканнер слов в модуле
	section  *sectionset.TSectionSet // разбивщик модуля на секции
	modules  *modules.TModules       // Набор модулей для компиляции
	filePath string                  // Имя главного файла для компиляции
	path     string                  // общий путь к файлу
}

// New -- возвращает новый *TOc
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
func (sf *TOc) Run() error {
	logrus.WithField("filePath", sf.filePath).Debugln("TOc.Run()")
	strSource, err := sf.readFile(sf.filePath)
	if err != nil {
		return fmt.Errorf("TOc.Run(): in read source file\n\t%w", err)
	}
	poolName := strings.Split(sf.filePath, "/")
	nameMod := types.AModule(poolName[len(poolName)-1])
	nameMod = nameMod[:len(nameMod)-3]
	sf.scanner.Scan(nameMod, strSource)
	// Разбить по секциям
	sf.section.Split(sf.scanner)
	nameMain := sf.section.Module().Name()
	sf.path = sf.filePath[:len(sf.filePath)-len(nameMain+".o7")]
	sf.checkModuleName(sf.filePath, nameMain)
	sf.modules.SetMain(nameMain)
	sf.modules.AddModule(sf.section.Module())
	sf.getImport()
	logrus.WithField("all modules", sf.modules.Len()).Debugln("Toc.Run()")
	sf.modules.ProcessConstant()
	return nil
}

// По требованию читает файл, возвращает содержимое
func (sf *TOc) readFile(filePath string) (strSource string, err error) {
	binSource, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("TOc.readFile(): error in read file\n\t%w", err)
	}
	strSource = string(binSource)
	return strSource, nil
}

// Проверяет имя модуля на соответствие имени файла
func (sf *TOc) checkModuleName(fileName string, moduleName types.AModule) {
	// Проверить, что имя главного модуля совпадает с имнем файла
	fileName = fileName[len(sf.path):]
	fileName = fileName[:len(fileName)-3]
	if fileName != string(moduleName) {
		logrus.Panicf("TOc.checkModuleName(): fileName(%v)!=moduleName(%v)\n", fileName, moduleName)
	}
}

func (sf *TOc) getImport() {
	// Получить остальные импорты модулей.
	modules := sf.section.GetImport()
	for _, module := range modules {
		if module == nil {
			logrus.Panicf("TOc.Run(): module==nil\n")
		}
		if err := sf.scanModule(module.Name()); err != nil {
			logrus.WithField("module", module.Name()).Panicf("TOc.Run(): in scam module\n")
		}
	}
}

// Готовит параметры под новый сканер
func (sf *TOc) scanModule(moduleName types.AModule) error {
	if sf.modules.IsExist(moduleName) {
		return nil
	}
	filePath := sf.path + string(moduleName) + ".o7"
	sf.scanner = scanner.New()
	sf.section = sectionset.New()
	strSource, err := sf.readFile(filePath)
	if err != nil {
		return fmt.Errorf("TOc.scanModule(): in read file\n\t%w", err)
	}
	sf.scanner.Scan(moduleName, strSource)
	sf.section.Split(sf.scanner)
	module := sf.section.Module()
	sf.checkModuleName(filePath, module.Name())
	sf.modules.AddModule(module)
	return nil
}
