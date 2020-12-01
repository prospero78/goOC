package module

/*
	Пакет предоставляет тип для вырезания слов модуля из пула слов.
	Назад возвращает все слова, что после модуля -- считается словами комментариев.
*/

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module/begin"
	"oc/internal/app/sectionset/module/consts"
	"oc/internal/app/sectionset/module/imports"
	"oc/internal/app/sectionset/module/keywords"
	"oc/internal/app/sectionset/module/otypes"
	"oc/internal/app/sectionset/module/procs"
	"oc/internal/app/sectionset/module/vars"
)

// TModule -- операции со словами модуля
type TModule struct {
	poolWord []*word.TWord       // пул слов модуля
	keywords *keywords.TKeywords // Пул допустимыз ключевых слов
	name     string              // Имя модуля
	imports  *imports.TImports   //  Секция импорта
	consts   *consts.TConsts     // Секция констант
	otypes   *otypes.TOtypes     // Секция типов
	vars     *vars.TVars         // Секция переменных
	procs    *procs.TProcedures  // Секция процедур
	begin    *begin.TBegin       // Секция BEGIN модуля
}

// New -- возвращает новый *TModule
func New() *TModule {
	return &TModule{
		poolWord: make([]*word.TWord, 0),
		keywords: keywords.New(),
		imports:  imports.New(),
		consts:   consts.New(),
		otypes:   otypes.New(),
		vars:     vars.New(),
		procs:    procs.New(),
		begin:    begin.New(),
	}
}

// Set -- вырезает все слова модуля, возвращает слова за модулем
func (sf *TModule) Set(poolWord []*word.TWord) (pool []*word.TWord, _len int) {
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
	sf.name = name.Word()
	delimeter := poolWord[2]
	if delimeter.Word() != ";" {
		log.Panicf("TModule.Set(): delimeter(%q) inname module bad\n", delimeter.Word())
	}
	poolWord = poolWord[3:]

	// Теперь перебрать все слова в модуле, лишнее вернуть
	for len(poolWord) >= 2 {
		word := poolWord[0]
		name := word.Word()
		if name == sf.name+"." {
			poolWord = poolWord[2:]
			log.Printf("TModule.Set(): name=%q word=%v\n", sf.name, len(sf.poolWord))
			return poolWord, len(sf.poolWord)
		}
		sf.poolWord = append(sf.poolWord, word)
		poolWord = poolWord[1:]
	}
	log.Panicf("TModule.Set(): not find end module\n")
	return nil, -1
}

// Split -- разделяет по требованию слова модуля по секциям
func (sf *TModule) Split() {
	poolWord := make([]*word.TWord, 0)
	poolWord = append(poolWord, sf.poolWord...)
	poolWord = sf.imports.Split(poolWord)
	poolWord = sf.consts.Split(poolWord)
	log.Printf("TModule.Split(): const=%v\n", sf.consts.Len())
	poolWord = sf.otypes.Split(poolWord)
	log.Printf("TModule.Split(): types=%v\n", sf.otypes.Len())
	poolWord = sf.vars.Split(poolWord)
	log.Printf("TModule.Split(): vars=%v\n", sf.vars.Len())
	poolWord = sf.procs.Split(poolWord)
	log.Printf("TModule.Split(): procs=%v\n", sf.procs.Len())
	poolWord = sf.begin.Split(poolWord)
	if len(poolWord)!=0{
		log.Panicf("TModule.Split(): after BEGIN poolWord len(%v)!=0\n", len(poolWord))
	}
}
