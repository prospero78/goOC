package module

/*
	Пакет предоставляет тип для вырезания слов модуля из пула слов.
	Назад возвращает все слова, что после модуля -- считается словами комментариев.
*/

import (
	"log"

	"github.com/prospero78/goOC/internal/app/sectionset/module/begin"
	"github.com/prospero78/goOC/internal/app/sectionset/module/consts"
	"github.com/prospero78/goOC/internal/app/sectionset/module/consts/srcconst"
	"github.com/prospero78/goOC/internal/app/sectionset/module/imports"
	"github.com/prospero78/goOC/internal/app/sectionset/module/imports/alias"
	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/app/sectionset/module/otypes"
	"github.com/prospero78/goOC/internal/app/sectionset/module/procs"
	"github.com/prospero78/goOC/internal/app/sectionset/module/vars"
	"github.com/prospero78/goOC/internal/types"
)

// TModule -- операции со словами модуля
type TModule struct {
	listWord []types.IWord      // пул слов модуля
	keywords types.IKeywords    // Пул допустимыз ключевых слов
	name     types.AModule      // Имя модуля
	imports  *imports.TImports  //  Секция импорта
	consts   *consts.TConsts    // Секция констант
	otypes   *otypes.TOtypes    // Секция типов
	vars     *vars.TVars        // Секция переменных
	procs    *procs.TProcedures // Секция процедур
	begin    *begin.TBegin      // Секция BEGIN модуля
}

// New -- возвращает новый *TModule
func New() *TModule {
	return &TModule{
		listWord: make([]types.IWord, 0),
		keywords: keywords.GetKeys(),
		imports:  imports.New(),
		consts:   consts.New(),
		otypes:   otypes.New(),
		vars:     vars.New(),
		procs:    procs.New(),
		begin:    begin.New(),
	}
}

// Set -- вырезает все слова модуля, возвращает слова за модулем
func (sf *TModule) Set(poolWord []types.IWord) (pool []types.IWord, _len int) {
	// Проверить наличие ключевого слова МОДУЛЬ, имени модуля и разделителя
	module := poolWord[0]
	if !sf.keywords.IsKey("MODULE", module.Word()) {
		log.Panicf("TModule.Set(): word(%q)!='MODULE'\n", module.Word())
	}
	name := poolWord[1]
	if !name.IsName() {
		log.Panicf("TModule.Set(): name(%q) not strong'\n", module.Word())
	}
	if sf.name != "" {
		log.Panicf("TModule.Set(): name(%q) already set\n", sf.name)
	}
	sf.name = types.AModule(name.Word())
	delimeter := poolWord[2]
	if delimeter.Word() != ";" {
		log.Panicf("TModule.Set(): delimeter(%q) inname module bad\n", delimeter.Word())
	}
	poolWord = poolWord[3:]

	// Теперь перебрать все слова в модуле, лишнее вернуть
	for len(poolWord) >= 2 {
		word := poolWord[0]
		name := types.AModule(word.Word())
		if name == sf.name+"." {
			poolWord = poolWord[2:]
			// log.Printf("TModule.Set(): name=%q word=%v\n", sf.name, len(sf.poolWord))
			return poolWord, len(sf.listWord)
		}
		sf.listWord = append(sf.listWord, word)
		poolWord = poolWord[1:]
	}
	log.Panicf("TModule.Set(): not find end module\n")
	return nil, -1
}

// Split -- разделяет по требованию слова модуля по секциям
func (sf *TModule) Split() {
	listWord := make([]types.IWord, 0)
	listWord = append(listWord, sf.listWord...)
	listWord = sf.imports.Split(listWord)
	log.Printf("TModule.Split() %q imports=%v", sf.name, sf.imports.Len())
	listWord = sf.consts.Split(listWord)
	// log.Printf("TModule.Split(): const=%v\n", sf.consts.Len())
	listWord = sf.otypes.Split(listWord)
	// log.Printf("TModule.Split(): types=%v\n", sf.otypes.Len())
	listWord = sf.vars.Split(listWord)
	// log.Printf("TModule.Split(): vars=%v\n", sf.vars.Len())
	listWord = sf.procs.Split(listWord)
	// log.Printf("TModule.Split(): procs=%v\n", sf.procs.Len())
	listWord = sf.begin.Split(listWord)
	if len(listWord) != 0 {
		log.Panicf("TModule.Split(): after BEGIN poolWord len(%v)!=0\n", len(listWord))
	}
}

// Name -- возвращает имя модуля после сканирования
func (sf *TModule) Name() types.AModule {
	return sf.name
}

// GetImport -- возвращает модули для импорта
func (sf *TModule) GetImport() []*alias.TAlias {
	return sf.imports.Imports()
}

// GetConst -- возвращает пул констант модуля
func (sf *TModule) GetConst() []*srcconst.TConst {
	return sf.consts.Get()
}
