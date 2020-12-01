package alias

import "log"

/*
	Пакет предоставляет тип для хранения имени модуля при импорте и его алиаса.
*/

// TAlias -- операции с именем и алиасом модуля
type TAlias struct {
	alias string // Алиас модуля
	name  string // Имя модуля
}

// New -- возвращает новый *TAlias
func New(name, alias string) *TAlias {
	{ // Предусловия
		if name == "" {
			log.Panicf("alias.go/New(): name==''\n")
		}
		if alias == "" {
			alias = name
		}

	}
	return &TAlias{
		alias: alias,
		name:  name,
	}
}

// Name -- возвращает имя модуля
func (sf *TAlias) Name() string {
	return sf.name
}

// Alias -- возвращает алиас модуля
func (sf *TAlias) Alias() string {
	return sf.alias
}
