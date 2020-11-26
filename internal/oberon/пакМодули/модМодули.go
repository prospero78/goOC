package пакМодули

/*
	Модуль предоставляет объект одиночку для учёта всех известных модулей в системе
*/

import (
	мСек "../пакСекция"
	мФмт "fmt"
)

//СМодульИмя -- специальный строковый тип для хранения имени модуля
type СМодульИмя string

//ИМодуль -- интерфейс интерфейса
type ИМодуль interface {
	мСек.ИСекция
	МодульИмя() СМодульИмя
	Обработать(мИф.СИсхФайл) error
}

//ИМодули -- интерфейс к глобальному объекту модулей
type ИМодули interface {
	МодульДоб(пСлово IWord, путьМодуль UWord) error //добавляет новый модуль в глобальный справочник модулей
}

//Тип глобального объекта экспортироваться не будет
type тМодули struct {
	спрМодули map[int]мИнт.ИМодульОпис
}

var (
	//Модули -- Объект управляет модулями при компиляции
	Модули мИнт.ИМодули
)

//Возвращает ссылку на новый и единственный тМодули
func модулиНов() (модули *тМодули, ош error) {
	if Модули != nil {
		return nil, мФмт.Errorf("модулиНов(): глобальный объект модулей уже существует\n")
	}
	модули = &тМодули{}
	if модули == nil {
		return nil, мФмт.Errorf("модулиНов(): нет памяти для глобального объекта модулей?\n")
	}
	модули.спрМодули = make(map[int]мИнт.ИМодульОпис)
	return модули, nil
}

//МодульДоб -- добавляет новый модуль в список
func (сам *тМодули) МодульДоб(пСлово мИнт.IWord, пПуть мИнт.UWord) (ош error) {
	if пСлово == nil {
		return мФмт.Errorf("тМодули.МодульДоб(): пСлово не может быть nil\n")
	}
	бЕсть := false //Есть ли модуль с таким именем в списке
	{              //1. Получить полный список хранимых имён модулей
		for _, модуль := range сам.спрМодули {
			if модуль.Имя() == мИнт.СМодульИмя(пСлово.Слово()) {
				бЕсть = true
				break
			}
		}
	}
	if бЕсть {
		return nil
	}
	//2. Модуля такого нет, создадим новый описатель
	модуль, ош := _МодульОписНов(пСлово, пПуть)
	if ош != nil {
		return мФмт.Errorf("тМодули.МодульДоб(): ошибка при создании описателя модуля\n\t%v", ош)
	}
	//3. Добавляем новый описатель модуля в общий справочник
	сам.спрМодули[len(сам.спрМодули)] = модуль
	мФмт.Printf("\tДобавлен модуль [%v:=%v]\n", пСлово.Слово(), пПуть)
	return nil
}

func init() {
	var ош error
	if Модули, ош = модулиНов(); ош != nil {
		panic(мФмт.Sprintf("модМодули.init(): ошибка при создании глобального объекта модулей\n\t%v", ош))
	}
}
