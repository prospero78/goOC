package keyword

/*
	Модуль предоставляет специальный тип для обработки ключевого(предопределённого) слова.
	У одного ключа может быть несколько синонимов, но как минимум один
*/

import (
	мФмт "fmt"
)

//СКлюч -- специальный строковый тип для хранения ключевого (предопределённого) слова
type СКлюч string

//ТКлюч -- предоставляет обработку ключевых (предопределённых) слов
type ТКлюч struct {
	пул map[int]СКлюч //Синонимы ключевого слова
}

//КлючНов -- возвращает указатель на новое ключевое слово
func КлючНов(пКлюч СКлюч) (ключ *ТКлюч) {
	_ключ := ТКлюч{
		пул: make(map[int]СКлюч),
	}
	_ключ.Доб(пКлюч)
	return &_ключ
}

//Доб -- добавляет синоним в набор ключевых слов
func (сам *ТКлюч) Доб(пКлюч СКлюч) {
	if пКлюч == "" {
		panic(мФмт.Errorf("ТКлюч.Доб(): пКлюч не может быть пустым\n"))
	}
	//Убедиться, что такого ключа нет
	for _, ключ := range сам.пул {
		if ключ == пКлюч {
			return
		}
	}
	сам.пул[len(сам.пул)] = пКлюч
}

//ЕслиСовпал -- проверяет совпадение ключевого слова с хранимым набором
func (сам *ТКлюч) ЕслиСовпал(пКлюч СКлюч) bool {
	if пКлюч == "" {
		panic(мФмт.Errorf("ТКлюч.ЕслиСовпал(): пКлюч не может быть пустым\n"))
	}
	//Перебрать в цикле ключевое слово
	for _, ключ := range сам.пул {
		if пКлюч == ключ {
			return true
		}
	}
	return false
}

//Синонимы -- возвращает синонимы ключа
func (сам *ТКлюч) Синонимы() map[int]СКлюч {
	рез := make(map[int]СКлюч)
	for адр, ключ := range сам.пул {
		рез[адр] = ключ
	}
	return рез
}