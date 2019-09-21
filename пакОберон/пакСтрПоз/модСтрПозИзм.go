package пакСтрПоз

/*
	Модуль предоставляет тип для операций с позицией в строке исходника
*/

import (
	мФмт "fmt"
	мСинх "sync"
)

//ИСтрПозИзм -- интерфейс к изменяемой позиции в строке
type ИСтрПозИзм interface{
	ИСтрПозФикс
	Уст(ССтрПоз)error
	Доб()
	Сброс()
}

//тСтрПозИзм -- тип для операций с позицией в строке исходника
type тСтрПозИзм struct {
	блок мСинх.RWMutex
	знач ССтрПоз
	стрПоз string
}

//СтрПозИзмНов -- возвращает ссылку на ИСтрПозИзм
func СтрПозИзмНов(пПоз ССтрПоз) (поз ИСтрПозИзм, ош error) {
	_поз := тСтрПозИзм{}
	if ош =_поз.Уст(пПоз);ош!=nil{
		return nil, мФмт.Errorf("СтрПозИзмНов(): ОШИБКА в начальном присвоении позиции\n\t%v", ош)
	}
	return &_поз, nil
}

//Уст -- многоразовая функция установки
func (сам *тСтрПозИзм) Уст(пПоз ССтрПоз) error {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	if пПоз < 0 {
		return мФмт.Errorf("тСтрПозИзм.Уст(): ОШИБКА значение меньше (0), пПоз=[%v]\n", пПоз)
	}
	сам.знач = пПоз
	сам.стрПоз = мФмт.Sprint(пПоз)
	return nil
}

//Доб -- добавляет +1 к значению позиции строки
func (сам *тСтрПозИзм) Доб() {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	сам.знач++
	сам.стрПоз = мФмт.Sprint(сам.знач)
}

//Сброс -- сбрасывает значение позиции строки
func (сам *тСтрПозИзм) Сброс() {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	сам.знач = 0
	сам.стрПоз = "0"
}

//Знач -- возвращает значение позиции в строке
func (сам *тСтрПозИзм) Знач() ССтрПоз {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.знач
}

func (сам *тСтрПозИзм) String() string {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.стрПоз
}
