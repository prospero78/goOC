package пакИсхСтрока

/*
	Модуль предоставляет тип исходной строки.
	Можно прочитать, получить исхНомер, но нельзя изменить
*/
import (
	мСтр "../../пакПулИсхСтроки/пакИсхСтрока/пакИсхСтрокаСтр"
	мНом "../../coord/coordy"
	мФмт "fmt"
)

//ИИсхСтрока -- интерфейс к исходной строке
type ИИсхСтрока interface {
	Строка() мСтр.СИсхСтрокаСтр
	Номер() int

//тИсхСтрока -- тип для операций со строкой исходника
type тИсхСтрока struct {
	исхСтрока мСтр.ИИсхСтрокаСтр
	исхНомер  мНом.ИСтрНомер
}

//ИсхСтрокаНов -- возвращает ссылку но новый тИсхСтрока
func ИсхСтрокаНов(пНом мНом.ССтрНомер, пСтр мСтр.СИсхСтрокаСтр) (стр ИИсхСтрока, ош error) {
	_стр := &тИсхСтрока{}
	if _стр == nil {
		return nil, мФмт.Errorf("ИсхСтрокаНов(): нет памяти для исходной строки?\n")
	}
	if _стр.исхСтрока, ош = мСтр.ИсхСтрокаСтрНов(пСтр); ош != nil {
		return nil, мФмт.Errorf("ИсхСтрокаНов(): ОШИБКА при создании исхСтрокаения исходной строки\n\t%v", ош)
	}
	if _стр.исхНомер, ош = мНом.СтрНомерНов(); ош != nil {
		return nil, мФмт.Errorf("ИсхСтрокаНов(): ОШИБКА при создании номера исходной строки\n\t%v", ош)
	}
	if ош = _стр.исхНомер.ИнитУст(пНом); ош != nil {
		return nil, мФмт.Errorf("ИсхСтрокаНов(): ОШИБКА при установке номера исходной строки\n\t%v", ош)
	}
	return _стр, nil
}

//исхСтрока -- возвращает хранимое исхСтрокаение исходной строки
func (сам *тИсхСтрока) Строка() мСтр.СИсхСтрокаСтр {
	return сам.исхСтрока.Получ()
}

//исхНомер -- возвращает хранимое исхСтрокаение номера исходной строки
func (сам *тИсхСтрока) Номер() мНом.ССтрНомер {
	return сам.исхНомер.Получ()
}
