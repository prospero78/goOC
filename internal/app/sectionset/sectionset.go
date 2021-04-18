package sectionset

/*
	Пакет предоставляет тип для разделения слов на набор секций.
	Виды секций:
		- комментарии (сборная солянка)
		- модуль (всё что относится к модулю)
		- импорт
		- константы
		- типы
		- переменные
		- процедуры
*/

import (
	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/app/sectionset/comment"
	"github.com/prospero78/goOC/internal/app/sectionset/module"
	"github.com/prospero78/goOC/internal/app/sectionset/module/imports/alias"
)

// IScan -- интерфейс к сканеру
type IScan interface {
	PoolWord() []*word.TWord
}

// TSectionSet -- операции с секциями
type TSectionSet struct {
	comment *comment.TComment // Слова комментариев
	module  *module.TModule   // Слова модуля
}

// New -- возвращает новый *TSectionSet
func New() *TSectionSet {
	return &TSectionSet{
		comment: comment.New(),
		module:  module.New(),
	}
}

// Split -- разделяет слова по секциям
func (sf *TSectionSet) Split(scanner IScan) {
	// log.Printf("TSectionSet()\n")
	poolWord := scanner.PoolWord()
	num := 0
	for _, word := range poolWord {
		if num != word.NumStr() {
			num = word.NumStr()
		}
	}
	poolWord = sf.comment.Set(poolWord)
	poolWord, _ = sf.module.Set(poolWord) // Дополнительные комментарии за концом файла
	_ = sf.comment.AddPoolWord(poolWord)
	// log.Printf("TSectionSet.Split(): comment+module=%v", numWordMod+numWordCom)

	// Теперь в модуле надо разделить слова по внутренним секциям
	sf.module.Split()
}

// ModuleName -- возвращает имя модуля после сканирования
func (sf *TSectionSet) ModuleName() string {
	return sf.module.Name()
}

// Module -- возвращает модуль после сканирования
func (sf *TSectionSet) Module() *module.TModule {
	return sf.module
}

// GetImport -- возвращает весь импорт модуля
func (sf *TSectionSet) GetImport() []*alias.TAlias {
	return sf.module.GetImport()
}
