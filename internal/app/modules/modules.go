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
	//   Для этого построить всё дерево импорта для рассматриваемого модуля из числа импортируемых.
	imports := module.GetImport()
	for _, modName := range imports {
		//poolName := make(map[string]bool)
		if modName.Name() == name {
			log.Panicf("TModules.AddModule(): for %q detected self import\n", name)
		}
		log.Printf("TModules.AddModule(): for %q find module %q\n", name, modName.Name())
	}
	sf.poolModule[name] = module
}

// Len -- возвращает число уникальных модулей
func (sf *TModules) Len() int {
	return len(sf.poolModule)
}
