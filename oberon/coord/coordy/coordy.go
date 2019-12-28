package coordy

/*
	Модуль предоставляет тип для операций с номером строки.
*/

import (
	мФмт "fmt"
	мСинх "sync"
)

//ТСтрНомер -- тип для операций с номером строки
type ТСтрНомер struct {
	знач     int
	стрНомер string
	блок     мСинх.RWMutex
}

//СтрНомерНов -- возвращает указатель на новый ТСтрНомер
func СтрНомерНов(пСтр int) (номер *ТСтрНомер) {
	_номер := ТСтрНомер{}
	_номер.Уст(пСтр)
	return &_номер
}

//Уст -- потокобезопасная установка значения номера строки
func (сам *ТСтрНомер) Уст(пНомер int) {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	if пНомер <= 0 {
		panic(мФмт.Errorf("ТСтрНомер.Уст(): ОШИБКА значение меньше 1, пНомер=[%v]\n", пНомер))
	}
	сам.знач = пНомер
	сам.стрНомер = мФмт.Sprintf("%v", пНомер)

}

//Получ -- возвращает хранимое значение номера строки
func (сам *ТСтрНомер) Получ() int {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.знач
}

func (сам *ТСтрНомер) String() string {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.стрНомер
}

//Доб -- добавляет значение номера строки
func (сам *ТСтрНомер) Доб() {
	сам.Уст(сам.знач + 1)
}

//Сброс -- сбрасывает значение номера строки
func (сам *ТСтрНомер) Сброс() {
	сам.Уст(1)
}
