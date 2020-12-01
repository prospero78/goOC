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
