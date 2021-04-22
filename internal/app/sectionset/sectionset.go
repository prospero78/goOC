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
	"github.com/prospero78/goOC/internal/app/sectionset/comment"
	"github.com/prospero78/goOC/internal/app/sectionset/module"
	"github.com/prospero78/goOC/internal/app/sectionset/module/imports/alias"
	"github.com/prospero78/goOC/internal/types"
)

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
func (sf *TSectionSet) Split(scanner types.IScanner) {
	// log.Printf("TSectionSet()\n")
	listWord := scanner.ListWord()
	num := types.ANumStr(0)
	for _, word := range listWord {
		if num != word.NumStr() {
			num = word.NumStr()
		}
	}
	listWord = sf.comment.Set(listWord)
	listWord, _ = sf.module.Set(listWord) // Дополнительные комментарии за концом файла
	_ = sf.comment.AddPoolWord(listWord)
	// log.Printf("TSectionSet.Split(): comment+module=%v", numWordMod+numWordCom)

	// Теперь в модуле надо разделить слова по внутренним секциям
	sf.module.Split()
}

// ModuleName -- возвращает имя модуля после сканирования
func (sf *TSectionSet) ModuleName() types.AModule {
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
