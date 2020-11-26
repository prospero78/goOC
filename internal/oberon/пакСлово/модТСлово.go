package пакСлово

/*
	Предоставляет тип слова для построения AST.
	Слово -- кусочек текста в исходнике.
	Обладает несколькими свойствами:
	1. Группа литер (само слово, может быть из одной литеры)
	2. Координаты.
	3. Строка, в которой он находится.
*/

import (
	мФмт "fmt"
	"oc/internal/types"

	мКоорд "../../пакОберон/пакКоорд"
	мИс "../пакПулИсхСтроки/пакИсхСтрока"
	мИсс "../пакПулИсхСтроки/пакИсхСтрока/пакИсхСтрокаСтр"
	мРод "../пакСлово/пакСловоРод"
)

//тСлово -- тип хранящий слово со всеми необходимыми атрибутами
type тСлово struct {
	стрИсх   types.UStringSource // Строка исходника
	стрСлово types.UWord         // Строка слова
	род      мРод.IWordРод       // род слова
	Коорд    types.ICoord        // Координаты слова в исходном тексте
	номер    types.UWordNum      // Текущий номер слова
	лит      types.ILit          //текущая литера
}

var словВсего = types.UWordNum(0)

//СловоНов -- возвращает ссылку на новый IWord
func СловоНов(пКоорд мКоорд.ИКоорд, пСлово types.UWord, пСтрИсх мИс.ИИсхСтрока) (слово types.IWord, ош error) {
	if len(пСлово) == 0 {
		return nil, мФмт.Errorf("СловоНовое(): строка имеет длину 0\n")
	}
	_слово := &тСлово{
		стрСлово: пСлово,
		стрИсх:   пСтрИсх,
		Коорд:    пКоорд,
		номер:    словВсего,
	}

	if _слово.лит, ош = мЛит.ЛитераНов(); ош != nil {
		return nil, мФмт.Errorf("СловоНовое(): ERROR при создании объекта литеры\n\t%v", ош)
	}
	if _слово.род, ош = мРод.СловоРодНов(мРод.UWordРод(_слово.Слово())); ош != nil {
		return nil, мФмт.Errorf("СловоНовое(): ERROR при создании рода слова\n\t%v", ош)
	}
	словВсего++
	return _слово, nil
}

//Род -- возвращает род слова
func (сам *тСлово) Род() мРод.СРод {
	return сам.род.Получ()
}

//Номер -- возвращает номер слова в исходнике
func (сам *тСлово) Номер() types.UWordNum {
	return сам.номер
}

// ЕслиПерваяБуква -- проверяет наличие первой буквы в строке
func (сам *тСлово) ЕслиПерваяБуква(пСтрока types.UWord) (бРез bool, ош error) {
	руна := []rune(пСтрока)
	стрЛит := мЛит.ULit(руна[0])
	if ош = сам.лит.Уст(стрЛит); ош != nil {
		return false, мФмт.Errorf("тСлово.ЕслиПерваяБуква(): ERROR в установке литеры\n\t%v", ош)
	}
	//пакКонс.Конс.Печать("пакСлово.ЕслиПерваяБуква(): лит=" + стрЛит)
	if сам.лит.IsLetter() {
		return true, nil
	}
	return false, nil
}

//ЕслиЧисло -- проверяет, что слово число
func (сам *тСлово) ЕслиЧисло(пСлово types.UWord) (бРез bool, ош error) {
	for _, лит := range пСлово {
		if ош = сам.лит.Уст(мЛит.ULit(лит)); ош != nil {
			return false, мФмт.Errorf("тСлово.ЕслиЧисло(): ERROR в установке литеры\n\t%v", ош)
		}
		if !сам.лит.IsDigit() {
			return false, nil
		}
	}
	return true, nil
}

// Проверяет, что строка не находится в ключевых словах
func (сам *тСлово) _ЕслиКлючевоеСлово() (бРез bool) {
	бРез = false // По умолчанию -- имя разрешено
	for _, группа := range КлючСлово {
		for _, ключевое := range группа {
			if сам.стрСлово == ключевое {
				бРез = true
				return бРез
			}
		}
	}
	return бРез
}

// Проверяет, что литера не находится в списке запрещённых (для имён сущностей)
func _ЕслиЛитЗапрещена(пЛит rune) (бРез bool) {
	//стрЗапрет:=[]rune("")
	стрЗапрет := []rune("\"~`!@$%^&*()-=+{}[]|\\<,>?/№;:\t\n'\r ")
	for _, лит := range стрЗапрет {
		if пЛит == лит {
			бРез = true
			return бРез
		}
	}
	return бРез
}

// IsName -- проверяет, что строка обладает строгим соответствием, чтобы быть именем сущности
func (сам *тСлово) IsName() bool {
	// имя сущности должно начинаться либо с "_", либо с буквы
	строка := []rune(сам.стрСлово)
	лит := мЛит.ULit(строка[0])
	if ош := сам.лит.Уст(лит); ош != nil {
		panic(мФмт.Sprintf("тСлово.IsName(): ERROR при присвоении литеры\n\t%v", ош))
	}
	if лит == "_" || сам.лит.IsLetter() {
		if лит != "_" && сам._ЕслиКлючевоеСлово() {
			return false
		}
		for _, лит := range строка {
			// Точка в имени -- допустимо, но здесь её не будет.
			if _ЕслиЛитЗапрещена(лит) {
				//
				return false
			}
		}
	}
	return true
}

//Слово -- возвращает слово, которое хранит тип
func (сам *тСлово) Слово() types.UWord {
	return сам.стрСлово
}

// Строка -- возвращает строку исходника, содержащую подстроку
func (сам *тСлово) Строка() мИсс.СИсхСтрокаСтр {
	return сам.стрИсх.Строка()
}

func (сам *тСлово) String() string {
	стрВых := мФмт.Sprintf("%v: %v [%v]", сам.Коорд, сам.стрИсх, сам.стрСлово)
	return стрВых
}
