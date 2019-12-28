package пакПулИсхСтроки

/*
	Пакет предоставляет тип, который хранит построчно исходник.
	К строке можно обратиться по номеру, или полностью получить пул строк
*/

import (
	мИф "../пакИсходникФайл"
	мФи "../пакИсходникФайл"
	мКонс "../пакКонсоль"
	мИс "../пакПулИсхСтроки/пакИсхСтрока"
	мИсс "../пакПулИсхСтроки/пакИсхСтрока/пакИсхСтрокаСтр"
	мНс "../пакСтрНомер"
	мФмт "fmt"
	мСтр "strings"
	мСинх "sync"
)

//ИПулИсхСтроки -- интерфейс для исходника разбитого построчно
type ИПулИсхСтроки interface {
	НаСтрокиРазбить(мИф.СИсхТекст) error
	Строка(мНс.ССтрНомер) (мИс.ИИсхСтрока, error)
	СтрокиВсе() map[мНс.ССтрНомер]мИс.ИИсхСтрока
}

//тПулИсхСтроки -- тип хранит список исходных строк
type тПулИсхСтроки struct {
	пулСтроки map[мНс.ССтрНомер]мИс.ИИсхСтрока
	блок      мСинх.RWMutex
}

//ПулИсхСтрокиНов -- возвращает новую ссылку на ИПулИсхСтроки
func ПулИсхСтрокиНов() (исх ИПулИсхСтроки) {
	_исх := тПулИсхСтроки{
		пулСтроки : make(map[мНс.ССтрНомер]мИс.ИИсхСтрока),
	}
	return &_исх
}

//НаСтрокиРазбить -- разбивает на строки содержимое строки
func (сам *тПулИсхСтроки) НаСтрокиРазбить(пИсх мФи.СИсхТекст) (ош error) {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	_ПулДоб := func(пНомер int, пСтр string) (ош error) {
		исхСтрока, ош := мИс.ИсхСтрокаНов(мНс.ССтрНомер(пНомер), мИсс.СИсхСтрокаСтр(пСтр))
		if ош != nil {
			return мФмт.Errorf("тИсхСтрока.НаСтрокиРазбить()._ПулДоб(): ОШИБКА при создании исходной строки\n\t%v", ош)
		}
		сам.пулСтроки[мНс.ССтрНомер(len(сам.пулСтроки))] = исхСтрока
		//сам.конс.Печать(пакФмт.Sprintf("%v: %v", итер, сам.пулСтроки[итер]))
		return nil
	}
	мКонс.Конс.Отладить("тПулИсхСтроки.НаСтроки_Разбить()")
	пулСтроки := мСтр.Split(string(пИсх), "\n")

	for ном, стр := range пулСтроки {
		if ош = _ПулДоб(ном+1, стр); ош != nil {
			return мФмт.Errorf("тПулИсхСтроки.НаСтрокиРазбить(): ОШИБКА при разбиении на строки\n\t%v", ош)
		}
	}
	return nil
}

//Строка -- возвращает строку по указанному номеру
func (сам *тПулИсхСтроки) Строка(пНомер мНс.ССтрНомер) (стр мИс.ИИсхСтрока, ош error) {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	for _, строка := range сам.пулСтроки {
		if строка.Номер() == пНомер {
			return строка, nil
		}
	}
	return nil, мФмт.Errorf("тПулИсхСтроки.Строка(): строки с номером [%v] не существует\n", пНомер)
}

func (сам *тПулИсхСтроки) СтрокиВсе() map[мНс.ССтрНомер]мИс.ИИсхСтрока {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.пулСтроки
}
