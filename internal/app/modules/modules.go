package modules

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module"
	"oc/internal/app/sectionset/module/keywords"
)

/*
	Пакет предоставляет тип для списка всех используемых модулей.
	Также хранит отдельно имя главного модуля.
*/

// TModules -- операции с модулями для компиляции
type TModules struct {
	mainName   string                     // Имя главного модуля
	poolModule map[string]*module.TModule // Пул модулей для компиляции
	keywords   *keywords.TKeywords        // Ключевые слова
}

// New -- возвращает новый *TModules
func New() *TModules {
	return &TModules{
		poolModule: make(map[string]*module.TModule),
		keywords:   keywords.Keys,
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
		sf.processConstant(module)
	}
}

// Проверяет простой тип константы.
func (sf *TModules) checkSimpleTypeConstant(word *word.TWord) (res string) {
	switch {
	case word.IsInt():
		return "INTEGER"
	case word.IsReal():
		return "REAL"
	case word.IsString():
		return "ARRAY OF CHAR"
	case word.IsBool():
		return "BOOLEAN"
	default:
		log.Panicf("TModules.checkTypeConstant(): unknown type=%v for constante\n", word.Word())
	}
	return ""
}

// Вычисляет выражение в скобках
func (sf *TModules) calcExp(pool []*word.TWord) {
	word := pool[0]
	switch {
	case word.IsName(): // Надо найти объект с таким именем

	default:
		log.Panicf("TModules.calcExp(): unknown type word(%v)\n", word.Word())
	}
}

// Если у константы много слов -- надо проверить их все.
func (sf *TModules) checkComplexTypeConstant(pool []*word.TWord) {
	wordBeg := pool[0]
	pool = pool[1:]
	switch {
	case wordBeg.Word() == "(": // Начало выражения
		poolExpr := make([]*word.TWord, 0) // Набор слов для выражения
		for {
			if len(pool) == 0 {
				break
			}
			if pool[0].Word() == ")" {
				pool = pool[1:] // Отбросить закрывающую скобку
				break
			}
			poolExpr = append(poolExpr, pool[0])
			pool = pool[1:]
		}

		sf.calcExp(poolExpr)
	default:
		log.Panicf("TModules.checkComplexTypeConstant(): unknown type for word(%v)\n", wordBeg.Word())
	}
}

// Проверяет типы констант в отдельном модуле
func (sf *TModules) processConstant(module *module.TModule) {
	poolConst := module.GetConst()
	for _, rawConst := range poolConst {
		if rawConst == nil {
			log.Panicf("TModules.GetConst(): rawConst==nil\n")
		}
		pool := rawConst.GetWords()
		switch len(pool) {
		case 0: // У константы нет имени. Теоретически, это невозможно
			log.Panicf("TModules.processConstant(): const(%v.%v) not have type\n", module.Name(), rawConst.Name())
		case 1: // У константы есть единственный тип
			sf.checkSimpleTypeConstant(pool[0])
		default: // Возможно, сложное выражение
			sf.checkComplexTypeConstant(pool)
		}
	}
}
