package пакСтрПоз

/*
	Модуль предоставляет тип для операций с позицией в строке исходника
*/

import (
	мФмт "fmt"
	мСинх "sync"
)

//ССтрПоз -- специальный целочисленный тип к позиции в строке
type ССтрПоз int

//ТСтрПоз -- тип для операций с позицией в строке исходника
type ТСтрПоз struct {
	блок   мСинх.RWMutex
	знач   ССтрПоз
	стрПоз string
}

//СтрПозНов -- возвращает ссылку на ИСтрПозИзм
func СтрПозНов(пПоз ССтрПоз) (поз *ТСтрПоз, ош error) {
	_поз := ТСтрПоз{}
	if ош = _поз.Уст(пПоз); ош != nil {
		return nil, мФмт.Errorf("СтрПозНов(): ОШИБКА в начальном присвоении позиции\n\t%v", ош)
	}
	return &_поз, nil
}

//Уст -- многоразовая функция установки
func (сам *ТСтрПоз) Уст(пПоз ССтрПоз) error {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	if пПоз < 0 {
		return мФмт.Errorf("ТСтрПоз.Уст(): ОШИБКА значение меньше (0), пПоз=[%v]\n", пПоз)
	}
	сам.знач = пПоз
	сам.стрПоз = мФмт.Sprint(пПоз)
	return nil
}

//Доб -- добавляет +1 к значению позиции строки
func (сам *ТСтрПоз) Доб() {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	сам.знач++
	сам.стрПоз = мФмт.Sprint(сам.знач)
}

//Сброс -- сбрасывает значение позиции строки
func (сам *ТСтрПоз) Сброс() {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	сам.знач = 0
	сам.стрПоз = "0"
}

//Получ -- возвращает значение позиции в строке
func (сам *ТСтрПоз) Получ() ССтрПоз {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.знач
}

func (сам *ТСтрПоз) String() string {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.стрПоз
}
