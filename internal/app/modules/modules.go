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
	mainName   string            // Имя главного модуля
	poolModule []*module.TModule // Пул модулей для компиляции
}

// New -- возвращает новый *TModules
func New() *TModules {
	return &TModules{
		poolModule: make([]*module.TModule, 0),
	}
}

// SetMain -- устанавливает главный модуль программы
func (sf *TModules) SetMain(name string) {
	if name == "" {
		log.Panicf("TModules.SetMain(): name(%v)==''\n", name)
	}
	if sf.mainName != "" {
		log.Panicf("TModules.SetMain(): name(%v)!=''\n", name)
	}
	sf.mainName = name
}
