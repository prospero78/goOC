package pullsourcestr

/*
	Пакет предоставляет тип, который хранит построчно исходник.
	К строке можно обратиться по номеру, или полностью получить пул строк
*/

import (
	"fmt"
	мИс "oc/internal/app/module/scannerword/pullsourcestr/sourcestr"
	//мЛог "oc/internal/log"
	"oc/internal/log"
	мТип "oc/internal/types"
	мСтр "strings"
)

//TPoolSource -- тип хранит список исходных строк
type TPoolSource struct {
	poolSource map[int]мТип.ИСтрокаИсх
	log       *log.TLog
}

// New -- возвращает новый *ТПулИсхСтроки
func New(текстИсх мТип.СТекстИсх, режим int) (poolSource *TPoolSource) {
	poolSource := &TPoolSource{
		log:       мЛог.Нов("TPoolSource", режим),
		poolSource: make(map[int]мТип.ИСтрокаИсх),
	}
	poolSource._НаСтрокиРазбить(текстИсх)
	return poolSource
}

//НаСтрокиРазбить -- разбивает на строки содержимое строки
func (sf *TPoolSource) _НаСтрокиРазбить(пИсх мТип.СТекстИсх) {
	пулСтроки := мСтр.Split(string(пИсх), "\n")

	for ном, стр := range пулСтроки {
		исхСтрока, _ := мИс.Нов(мТип.ССтрокаНом(ном+1), мТип.ССтрокаИсх(стр))
		sf.poolSource[len(sf.poolSource)+1] = исхСтрока
	}
	sf.log.Отладка("_НаСтрокиРазбить", "всего строк: ", len(пулСтроки))
}

//Строка -- возвращает строку по указанному номеру
func (sf *TPoolSource) Строка(пНомер мТип.ССтрокаНом) (стр мТип.ИСтрокаИсх, ош error) {
	стр, ок := sf.poolSource[int(пНомер)]
	if !ок {
		return nil, fmt.Errorf("TPoolSource.Строка(): строки с номером [%v] не существует\n", пНомер)
	}
	return стр, nil
}

func (sf *TPoolSource) СтрокиВсе() map[int]мТип.ИСтрокаИсх {
	return sf.poolSource
}
