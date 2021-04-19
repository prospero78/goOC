package alias

import (
	"log"

	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип для хранения имени модуля при импорте и его алиаса.
*/

// TAlias -- операции с именем и алиасом модуля
type TAlias struct {
	alias types.AModule // Алиас модуля
	name  types.AModule // Имя модуля
}

// New -- возвращает новый *TAlias
func New(name, alias types.AModule) *TAlias {
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
func (sf *TAlias) Name() types.AModule {
	return sf.name
}

// Alias -- возвращает алиас модуля
func (sf *TAlias) Alias() types.AModule {
	return sf.alias
}
