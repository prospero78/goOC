// Package modules -- Пакет предоставляет тип для списка всех используемых модулей.
//  Также хранит отдельно имя главного модуля.
package modules

import (
	"log"

	"github.com/prospero78/goOC/internal/app/modules/calcconst"
	"github.com/prospero78/goOC/internal/app/modules/calcword"
	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/app/sectionset/module"
	"github.com/prospero78/goOC/internal/app/sectionset/module/consts/srcconst"
	"github.com/prospero78/goOC/internal/app/sectionset/module/consts/srcconst/constexpres"
	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/types"
)

// TModules -- операции с модулями для компиляции
type TModules struct {
	mainName    types.AModule                     // Имя главного модуля
	poolModule  map[types.AModule]*module.TModule // Пул модулей для компиляции
	keywords    types.IKeywords                   // Ключевые слова
	modCurrent  *module.TModule                   // Текущий обрабатываемый модуль
	consCurrent *srcconst.TConst                  // Текущая константа
	expCurrent  *constexpres.TConstExpression     // Текущее выражение для вычисления
	wordCurrent *word.TWord                       // Текущее слово для обработки

	stackConst    []*srcconst.TConst              // Стек для констант
	stackConstExp []*constexpres.TConstExpression // стек для выражений констант
	calcConst     *calcconst.TCalcConst           // Калькулятор констант
	calcWord      *calcword.TCalcWord             // Калькулятор отдельного слова
}

// New -- возвращает новый *TModules
func New() *TModules {
	return &TModules{
		poolModule: make(map[types.AModule]*module.TModule),
		keywords:   keywords.GetKeys(),
		stackConst: make([]*srcconst.TConst, 0),
		calcConst:  calcconst.New(),
		calcWord:   calcword.New(),
	}
}

// SetMain -- устанавливает главный модуль программы
func (sf *TModules) SetMain(name types.AModule) {
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
func (sf *TModules) checkImport(module *module.TModule, name types.AModule) {
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
func (sf *TModules) IsExist(modname types.AModule) bool {
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
			sf.calcConst.Calc(cons)
		}
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

// После обработки всех слов константы -- устанавливает её тип
func (sf *TModules) setConstType() {
	pool := sf.consCurrent.GetWords()
	switch len(pool) {
	case 0: // Нет слов у константы (теоретически такого быть не может)
		log.Panicf("TModules.setConstType(): const(%v.%v) not have type\n", sf.modCurrent.Name(), sf.consCurrent.Name())
	case 1: // Тип константы определяется единственным словом
		sf.consCurrent.SetType(pool[0].GetType())
	default: // Тип имеет выражение и его надо вычислить
		// exp := sf.consCurrent.GetExpres()
		// sf.exprConstCalc(exp)
		sf.stackConstExp = append(sf.stackConstExp, sf.expCurrent)
		sf.expCurrent = sf.consCurrent.GetExpres()
		listOld := sf.consCurrent.GetWords()
		listOld = listOld[1:] // Откинуть открывающую скобку
		for len(listOld) > 0 {
			word := listOld[0]
			sf.expCurrent.AddWord(word)
			listOld = listOld[1:]
			if word.Word() == ")" {
				break
			}
		}
		sf.exprConstCalc()
		// После передачи слов в выражение -- надо сформировать новый словарь слов
		listNew := make([]types.IWord, 0)
		listNew = append(listNew, sf.expCurrent.GetWord())
		listNew = append(listNew, listOld...)
		sf.consCurrent.SetPoolWord(listNew)
		sf.exprConstCalc()
		sf.expCurrent = sf.stackConstExp[len(sf.stackConstExp)-1]
	}
}

// Вычисляет выражения в константых
func (sf *TModules) exprConstCalc() {
	pool := sf.expCurrent.GetWords()
	switch len(pool) {
	case 0: // Теоретически такое невозможно
		log.Panicf("TModules.exprConstCalc(): const(%v.%v) not have word in expression\n", sf.modCurrent.Name(), sf.consCurrent.Name())
	case 1: // Такое теоретически возможно (если одно слово в скобках)
		word := pool[0]
		sf.calcWord.RecognizeType(word)
		sf.expCurrent.SetType(sf.wordCurrent.GetType())
	}
}
