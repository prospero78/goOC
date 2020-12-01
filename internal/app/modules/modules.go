package modules

import (
	"log"
	"oc/internal/app/sectionset/module"
)

/*
	Пакет предоставляет тип для списка всех используемых модулей.
	Также хранит отдельно имя главного модуля.
*/

// TModules -- операции с модулями для компиляции
type TModules struct {
	mainName   string                     // Имя главного модуля
	poolModule map[string]*module.TModule // Пул модулей для компиляции
}

// New -- возвращает новый *TModules
func New() *TModules {
	return &TModules{
		poolModule: make(map[string]*module.TModule),
	}
}

// SetMain -- устанавливает главный модуль программы
func (sf *TModules) SetMain(name string, module *module.TModule) {
	{ // Предусловия
		if name == "" {
			log.Panicf("TModules.SetMain(): name(%v)==''\n", name)
		}
		if module == nil {
			log.Panicf("TModules.SetMain(): module==nil\n")
		}
	}

	if sf.mainName != "" {
		log.Panicf("TModules.SetMain(): name(%v)!=''\n", name)
	}
	sf.mainName = name
	sf.poolModule[name] = module
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
