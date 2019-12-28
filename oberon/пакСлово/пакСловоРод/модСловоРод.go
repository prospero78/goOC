package пакСловоРод

/*
	Модуль предоставляет тип для хранения рода слова
*/

import (
	мФмт "fmt"
	мСинх "sync"
	//мСтр "strings"
	мЛит "../../liter"
)

//СРод -- специальный целочисленный тип для хранения рода слова
type СРод int

//ССловоРод -- специальный строковый тип для хранения слова, в котором надо распознать сам.знач
type ССловоРод string

//ИСловоРод - -интерфейс к операциям с родом слова
type ИСловоРод interface {
	Получ() СРод
}

type тСловоРод struct {
	знач СРод
	блок мСинх.RWMutex
}

//СловоРодНов -- возвращает ссылку на новый ИСловоРод
func СловоРодНов(пСлово ССловоРод) (род ИСловоРод, ош error) {
	_род := &тСловоРод{}
	if _род == nil {
		return nil, мФмт.Errorf("СловоРодНов(): нет памяти под новый сам.знач слова?\n")
	}
	if ош = _род._Уст(пСлово); ош != nil {
		return nil, мФмт.Errorf("СловоРодНов(): ОШИБКА при установке рода слова\n\t%v", ош)
	}
	return _род, nil
}

func (сам *тСловоРод) Получ() СРод {
	defer сам.блок.RUnlock()
	сам.блок.RLock()
	return сам.знач
}

func (сам *тСловоРод) _Уст(пСлово ССловоРод) (ош error) {
	defer сам.блок.Unlock()
	сам.блок.Lock()
	if пСлово == "" {
		return мФмт.Errorf("тСловоРод._Уст(): пСлово не может быть пустым\n")
	}
	switch пСлово {
	case ";":
		{
			сам.знач = КТочкаЗапятая
			return nil
		}
	case ",":
		{
			сам.знач = КЗапятая
			return nil
		}
	case "+":
		{
			сам.знач = КПлюс
			return nil
		}
	case "-":
		{
			сам.знач = КМинус
			return nil
		}
	case "/":
		{
			сам.знач = КДелить
			return nil
		}
	case "(":
		{
			сам.знач = КСкобкаОткрКругл
			return nil
		}
	case "(*":
		{
			сам.знач = ККомментНачать
			return nil
		}
	case ")":
		{
			сам.знач = КСкобкаЗакрКругл
			return nil
		}
	case "*)":
		{
			сам.знач = ККомментЗакончить
			return nil
		}
	case "*":
		{
			сам.знач = КУмножить
			return nil
		}
	case ":=":
		{
			сам.знач = КПрисвоить
			return nil
		}
	case ":":
		{
			сам.знач = КОпределить
			return nil
		}
	case "=":
		{
			сам.знач = КРавно
			return nil
		}
	case ".":
		{
			сам.знач = КТочка
			return nil
		}
	}
	{ //Проверка на строку
		слНач := пСлово[0]
		слКон := пСлово[len(пСлово)-1]
		if слНач == '"' && слКон == '"' {
			сам.знач = КСтрока
			return nil
		}
	}
	лит, ош := мЛит.ЛитераНов()
	if ош != nil {
		return мФмт.Errorf("тСловоРод._Уст(): ОШИБКА при создании литеры\n\t%v", ош)
	}
	{ //Проверка на имя
		литПрефикс := string([]rune(пСлово)[0])
		бПрефикс := литПрефикс == "_"
		if ош = лит.Уст(мЛит.СЛит(литПрефикс)); ош != nil {
			return мФмт.Errorf("тСловоРод._Уст(): ОШИБКА при установке префикса(%v)\n\t%v", литПрефикс, ош)
		}
		бБуква := лит.ЕслиБуква()
		if бПрефикс || бБуква {
			сам.знач = КИмя
			return nil
		}
	}
	{ //Проверка на цифру
		if лит.ЕслиЦифра() {
			сам.знач = КЧисло
			return nil
		}
	}
	return мФмт.Errorf("тСловоРод._Уст(): ОШИБКА не могу установить род строки, строка=" + string(пСлово))
}
