// Package modules -- Пакет предоставляет тип для списка всех используемых модулей.
//  Также хранит отдельно имя главного модуля.
package modules

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module"
	"oc/internal/app/sectionset/module/consts/srcconst"
	"oc/internal/app/sectionset/module/consts/srcconst/constexpres"
	"oc/internal/app/sectionset/module/keywords"
	"strings"
)

// TModules -- операции с модулями для компиляции
type TModules struct {
	mainName    string                        // Имя главного модуля
	poolModule  map[string]*module.TModule    // Пул модулей для компиляции
	keywords    *keywords.TKeywords           // Ключевые слова
	modCurrent  *module.TModule               // Текущий обрабатываемый модуль
	consCurrent *srcconst.TConst              // Текущая константа
	expCurrent  *constexpres.TConstExpression // Текущее выражение для вычисления
	wordCurrent *word.TWord                   // Текущее слово для обработки

	stackConst    []*srcconst.TConst              // Стек для констант
	stackConstExp []*constexpres.TConstExpression // стек для выражений констант
}

// New -- возвращает новый *TModules
func New() *TModules {
	return &TModules{
		poolModule: make(map[string]*module.TModule),
		keywords:   keywords.Keys,
		stackConst: make([]*srcconst.TConst, 0),
	}
}

// SetMain -- устанавливает главный модуль программы
func (sf *TModules) SetMain(name string) {
	{ // Предусловия
		if name == "" {
			log.Panicf("TModules.SetMain(): name(%v)==''\n", name)
		}
	}

	if sf.mainName != "" {
		log.Panicf("TModules.SetMain(): name(%v)!=''\n", name)
	}
	sf.mainName = name
}

// AddModule -- добавляет модуль в пул модулей, проверяет циклические ссылки
func (sf *TModules) AddModule(module *module.TModule) {
	if module == nil {
		log.Panicf("TModules.AddModule(): module==nil\n")
	}
	name := module.Name()
	if _, ok := sf.poolModule[name]; ok {
		return
	}
	// Проверить циклический импорт
	//   Не должен импортировать модули, которые импортируют его.
	//   Для этого проверим все модули, которые импортируют этот модуль
	imports := module.GetImport()
	for _, nameCheck := range imports {
		if name == nameCheck.Name() {
			log.Panicf("TModules.AddModule(): for module %q detected auto import\n", name)
		}
		modCheck, ok := sf.poolModule[nameCheck.Name()]
		if !ok {
			continue
		}
		sf.checkImport(modCheck, name)
	}

	sf.poolModule[name] = module
}

// Len -- возвращает число уникальных модулей
func (sf *TModules) Len() int {
	return len(sf.poolModule)
}

// Рекурсивно проверяет наличие циклического импорта для указанного имени модуля
func (sf *TModules) checkImport(module *module.TModule, name string) {
	imports := module.GetImport()
	// Найти имя модуля в импорте проверяемого модуля
	// Импортируемый модуль ещё может быть не зарегистрирован
	for _, modImp := range imports {
		if name == modImp.Name() {
			log.Panicf("TModules.AddModule(): for module %q detected cyclo import in module %q:%q\n", name, module.Name(), modImp.Name())
		}
	}
}

// IsExist -- возвращает признак существования модуля в реестре модулей
func (sf *TModules) IsExist(modname string) bool {
	_, ok := sf.poolModule[modname]
	return ok
}

// ProcessConstant -- обрабатывает все константы.
//   Берёт по порядку константы из модуля, проверяет их тип, если тип внешний --
//   пытается найти такой тип во внешнем модуле.
func (sf *TModules) ProcessConstant() {
	for _, module := range sf.poolModule {
		sf.modCurrent = module
		poolConst := module.GetConst()
		for _, cons := range poolConst {
			sf.consCurrent = cons
			sf.processConstant()
		}
	}
}

// Проверяет простой тип константы.
func (sf *TModules) checkTypeWordConst() {
	switch {
	case sf.wordCurrent.IsInt(): // Если целое
		sf.wordCurrent.SetType("INTEGER")
	case sf.wordCurrent.IsReal(): // Если вещественное
		sf.wordCurrent.SetType("REAL")
	case sf.wordCurrent.IsString(): // Если строка
		sf.wordCurrent.SetType("ARRAY OF CHAR")
	case sf.wordCurrent.IsBool(): // Если булево
		sf.wordCurrent.SetType("BOOLEAN")
	case sf.wordCurrent.Word() == "(": // Если начало выражения
		sf.wordCurrent.SetType("(")
	case sf.wordCurrent.Word() == ")": // Если конец выражения
		sf.wordCurrent.SetType(")")
	case sf.wordCurrent.IsName(): // Если присвоение из другой константы
		if sf.wordCurrent.GetType() == "" { // Если ещё это имя не встречалось -- найти его тип
			// Сохранить текущую константу в стеке
			sf.stackConst = append(sf.stackConst, sf.consCurrent)
			// Найти константу в списке констант модуля
			cons := sf.modCurrent.GetConst()
			sf.consCurrent = nil
			for _, con := range cons {
				if con.Name() == sf.wordCurrent.Word() {
					sf.consCurrent = con
					break
				}
			}
			if sf.consCurrent == nil { // убедиться, что такая константа существует
				log.Panicf("TModules.checkTypeWordConst(): unknown constante  %v.%v\n", sf.modCurrent.Name(), sf.wordCurrent.Word())
			}
			sf.processConstant()                                 // обработать новую константу
			sf.wordCurrent.SetType(sf.consCurrent.GetType())     // Установить тип текущей константы из обработанной
			sf.consCurrent = sf.stackConst[len(sf.stackConst)-1] // Восстановить текущую константу
			sf.stackConst = sf.stackConst[:len(sf.stackConst)-1]

		}
	case sf.wordCurrent.IsCompoundName(): // Имя состоит из нескольких частей
		poolName := strings.Split(sf.wordCurrent.Word(), ".")
		// Проверить, что "Модуль:имя"
		if len(poolName) == 2 {
			// Найти имя модуля
			modName := poolName[0]
			if _, ok := sf.poolModule[modName]; !ok { // К этому моменту все модули просканированы
				log.Panicf("TModules.checkTypeConstant(): unknown name module(%v) for constante %v.%v\n", sf.wordCurrent.Word(), sf.modCurrent.Name(), sf.wordCurrent.Word())
			}
		}
	case sf.wordCurrent.Word() == "+": // Операция "+"
		sf.wordCurrent.SetType("+")
	case sf.wordCurrent.Word() == "-": // Операция "-"
		sf.wordCurrent.SetType("-")
	case sf.wordCurrent.Word() == "/": // Операция РАЗДЕЛИТЬ
		sf.wordCurrent.SetType("/")
	case sf.wordCurrent.Word() == "*": // Операция "*"
		sf.wordCurrent.SetType("*")
	default:
		log.Panicf("TModules.checkTypeConstant(): unknown type for constante %v.%v\n", sf.modCurrent.Name(), sf.wordCurrent.Word())
	}
}

// Вычисляет выражение в скобках
func (sf *TModules) calcExp(pool []*word.TWord) {
	word := pool[0]
	switch {
	case word.IsName(): // Надо найти объект с таким именем
		cons := sf.findConstName(word)
		if len(cons.GetWords()) == 0 {
			log.Panicf("TModules.calcExp(): const(%v) not have type\n", cons.Name())
		}
		if cons.GetType() != sf.consCurrent.GetType() {
			log.Panicf("TModules.calcExp(): type cons.type(%v)!=consCurrent.type(%v)\n", cons.GetType(), sf.consCurrent.GetType())
		}
	default:
		log.Panicf("TModules.calcExp(): unknown type word(%v)\n", word.Word())
	}
}

// Ищет простое имя в текущем обрабатываемом модуле
func (sf *TModules) findConstName(name *word.TWord) *srcconst.TConst {
	for _, cons := range sf.modCurrent.GetConst() {
		if cons.Name() == name.Word() {
			return cons
		}
	}
	log.Panicf("TModules.findSimpleName(): not find constante %v.%v\n", sf.modCurrent.Name(), name.Word())
	return nil
}

// Проверяет типы констант в отдельном модуле
func (sf *TModules) processConstant() {
	pool := sf.consCurrent.GetWords()
	if len(pool) == 0 { // У константы нет имени. Теоретически, это невозможно
		log.Panicf("TModules.processConstant(): const(%v.%v) not have type\n", sf.modCurrent.Name(), sf.consCurrent.Name())
	}
	if sf.consCurrent.Name() == "\"цЯблоки\"" {
			log.Print("")
		}
	lenPool := len(pool)
	fnCheckWord := func() bool {
		adr := 0
		for {
			pool = sf.consCurrent.GetWords()
			if adr >= len(pool) {
				sf.setConstType()
				return false
			}
			sf.wordCurrent = pool[adr]
			adr++
			sf.checkTypeWordConst()
			_lenPool := len(pool)
			if _lenPool < lenPool {
				lenPool = _lenPool
				return true
			}
		}
	}

	if len(pool) == 1 {
		sf.wordCurrent = pool[0]
		sf.checkTypeWordConst()
		sf.setConstType()
		return
	}
	for fnCheckWord() {
	}
}

// После обработки всех слов константы -- устанавливает её тип
func (sf *TModules) setConstType() {
	pool := sf.consCurrent.GetWords()
	switch len(pool) {
	case 0: // Нет слов у константы (теоретически такого быть не может)
		log.Panicf("TModules.processConstant(): const(%v.%v) not have type\n", sf.modCurrent.Name(), sf.consCurrent.Name())
	case 1: // Тип константы определяется единственным словом
		sf.consCurrent.SetType(pool[0].GetType())
	default: // Тип имеет выражение и его надо вычислить
		//exp := sf.consCurrent.GetExpres()
		//sf.exprConstCalc(exp)
		sf.stackConstExp = append(sf.stackConstExp, sf.expCurrent)
		sf.expCurrent = sf.consCurrent.GetExpres()
		poolWord := sf.consCurrent.GetWords()
		poolWord = poolWord[1:] // Откинуть открывающую скобку
		for len(poolWord) > 0 {
			word := poolWord[0]
			sf.expCurrent.AddWord(word)
			poolWord = poolWord[1:]
			if word.Word() == ")" {
				break
			}
		}
		sf.exprConstCalc()
		// После передачи слов в выражение -- надо сформировать новый словарь слов
		poolNew := make([]*word.TWord, 0)
		poolNew = append(poolNew, sf.expCurrent.GetWord())
		poolNew = append(poolNew, poolWord...)
		sf.consCurrent.SetPoolWord(poolNew)
		sf.exprConstCalc()
		sf.expCurrent = sf.stackConstExp[len(sf.stackConstExp)-1]
	}
}

// Вычисляет выражения в константых
func (sf *TModules) exprConstCalc() {
	pool := sf.expCurrent.GetWords()
	switch len(pool) {
	case 0: // Теоретически такое невозможно
		log.Panicf("TModules.processConstant(): const(%v.%v) not have word in expression\n", sf.modCurrent.Name(), sf.consCurrent.Name())
	case 1: // Такое теоретически возможно (если одно слово в скобках)
		sf.wordCurrent = pool[0]
		sf.checkTypeWordConst()
		sf.expCurrent.SetType(sf.wordCurrent.GetType())
	}
}
