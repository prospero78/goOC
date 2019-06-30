package пакКлючи

/*
	Модул  предоставляет объект для хранения всех допустимых ключей в Обероне
*/

import (
	мФмт "fmt"
	мИнт "github.com/prospero78/goOC/пакОберон/пакИнтерфейсы"
)

//тКлючиВсе -- тип для обработки всех допустимых ключей в Обероне
type тКлючиВсе struct {
	спрКлючи map[мИнт.СКлюч]мИнт.ИКлюч
}

var (
	//Ключи -- объект для обработки всех допустимых ключей в Обероне
	Ключи *тКлючиВсе
)

func _КлючиВсеНов() {
	Ключи = &тКлючиВсе{}
	if Ключи == nil {
		panic(мФмт.Sprintf("_КлючиВсеНов(): нет памяти под все ключи?\n"))
	}
	Ключи.спрКлючи = make(map[мИнт.СКлюч]мИнт.ИКлюч)

}

func (сам *тКлючиВсе) _Доб(пКлюч мИнт.СКлюч, пКлючи []мИнт.СКлюч) {
	_Паника := func(пКлюч мИнт.СКлюч, ош error) {
		if ош != nil {
			panic(мФмт.Sprintf("тКлючиВсе._Доб(): ошибка добавления ключа \"%v\"\n\t%v", пКлюч, ош))
		}
	}
	if len(пКлючи) == 0 {
		panic(мФмт.Sprintf("тКлючиВсе._Доб(): пКлюч не может иметь нулевую длину\n"))
	}
	ключ, ош := КлючНов(пКлючи[0])
	_Паника(пКлючи[0], ош)
	for адр := 1; адр < len(пКлючи); адр++ {
		ош = ключ.Доб(пКлючи[адр])
		_Паника(пКлючи[адр], ош)
	}
	сам.спрКлючи[пКлюч] = ключ
}

/*СекцияТипПров -- проверяет, что переданный ключ реально является секцией
  Допустимые ключи:
  MODULE -- псевдосекция для модуля
  COMMENTDEF -- псевдосекция для комментариев
  IMPORT
  TYPE
  CONST
  VAR
  PROCDEF -- псевдосекция для процедур
  BEGIN -- секция инициализации модуля
*/
func (сам *тКлючиВсе) СекцияТипПров(пТип мИнт.СКлюч) (ош error) {
	if пТип == "" {
		return мФмт.Errorf("тКлючиВсе.СекцияТипПров(): пКлюч не может иметь нулевую длину\n")
	}
	//Проверить, а существует ли такой ключ секции вообще
	for _, типСекции := range []string{"MODULE", "COMMENTDEF", "IMPORT", "TYPE", "CONST", "VAR", "PROCDEF", "BEGIN"} {
		//Проверить есть ли в базе синонимов такой тип секции
		синонимы := сам.спрКлючи[мИнт.СКлюч(типСекции)]
		for _, тип := range синонимы.Синонимы() {
			if тип == пТип {
				return nil
			}
		}
	}
	return мФмт.Errorf("тКлючиВсе.СекцияТипПров(): в Обероне не может быть такого типа секции, пТип=[%v]\n", пТип)
}

func (сам *тКлючиВсе) Проверить(пКлюч, пСлово мИнт.СКлюч) (ок bool, ош error) {
	if пКлюч == "" {
		return false, мФмт.Errorf("тКлючиВсе.Проверить(): пКлюч не может иметь нулевую длину\n")
	}
	//Проверить, а существует ли такой ключ вообще
	бЕсть := false
	for ключ := range сам.спрКлючи {
		if ключ == пКлюч {
			бЕсть = true
			break
		}
	}
	if !бЕсть {
		return false, мФмт.Errorf("тКлючиВсе.Проверить(): в Обероне не может быть такого ключа, пКлюч=[%v]\n", пКлюч)
	}
	//Проверить есть ли в базе синонимы
	синонимы := сам.спрКлючи[пКлюч]
	for _, ключ := range синонимы.Синонимы() {
		if ключ == пСлово {
			return true, nil
		}
	}

	return false, nil
}

func init() {
	_КлючиВсеНов()
	Ключи._Доб("MODULE", []мИнт.СКлюч{"MODULE", "МОДУЛЬ"})
	Ключи._Доб("COMMENTDEF", []мИнт.СКлюч{"COMMENTDEF", "КОММЕНТАРИИ", "КОММЕНТ"})
	Ключи._Доб("END", []мИнт.СКлюч{"END", "КОНЕЦ"})
	Ключи._Доб("IMPORT", []мИнт.СКлюч{"IMPORT", "ИМПОРТ"})
	Ключи._Доб("TYPE", []мИнт.СКлюч{"TYPE", "ТИПЫ"})
	Ключи._Доб("CONST", []мИнт.СКлюч{"CONST", "КОНСТ"})
	Ключи._Доб("VAR", []мИнт.СКлюч{"VAR", "ПЕРЕМ"})
	Ключи._Доб("PROCDEF", []мИнт.СКлюч{"PROCDEF", "ПРОЦОПР"})
	Ключи._Доб("BEGIN", []мИнт.СКлюч{"BEGIN", "НАЧАЛО", "НАЧ"})
	Ключи._Доб(";", []мИнт.СКлюч{";"})
	Ключи._Доб(".", []мИнт.СКлюч{"."})
	Ключи._Доб(",", []мИнт.СКлюч{","})
	Ключи._Доб(":=", []мИнт.СКлюч{":="})
}
