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
	poolSource map[int]мТип.ISourceString
	log        *log.TLog
}

// New -- возвращает новый *ТПулИсхСтроки
func New(текстИсх мТип.UTextSource, режим int) (poolSource *TPoolSource) {
	poolSource := &TPoolSource{
		log:        мЛог.Нов("TPoolSource", режим),
		poolSource: make(map[int]мТип.ISourceString),
	}
	poolSource._НаСтрокиРазбить(текстИсх)
	return poolSource
}

// toStringSplit -- разбивает на строки содержимое строки
func (sf *TPoolSource) toStringSplit(txtSource мТип.UTextSource) {
	poolString := мСтр.Split(string(txtSource), "\n")

	for adr, str := range poolString {
		strSource, _ := мИс.Нов(мТип.UStringNum(adr+1), мТип.UStringSource(str))
		sf.poolSource[len(sf.poolSource)+1] = strSource
	}
	sf.log.Debugf("toStringSplit", "всего строк: ", len(poolString))
}

// GetString -- возвращает строку по указанному номеру
func (sf *TPoolSource) GetString(пНомер мТип.UStringNum) (strSource мТип.ISourceString, err error) {
	strSource, ок := sf.poolSource[int(пНомер)]
	if !ок {
		return nil, fmt.Errorf("TPoolSource.GetString(): строки с номером [%v] не существует", пНомер)
	}
	return strSource, nil
}

// GetPool -- возвращает пул строк
func (sf *TPoolSource) GetPool() map[int]мТип.ISourceString {
	return sf.poolSource
}
