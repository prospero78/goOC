package pullkeyword

/*
	Модуль предоставляет пул ключевых слов. Это разделяемый объект в единственном числе.
*/

import (
	"oc/internal/app/module/scannerword/word/poolkeyword/keyword"
	мТип "oc/internal/types"
)

//тПулКлючСлова -- операции с пулом ключевых слов
type тПулКлючСлова struct {
	пулСлова map[мТип.UWordKey]*keyword.ТКлюч
}

var (
	//ПулКлюч -- пул всех доступных ключей, синглетон
	ПулКлюч *тПулКлючСлова
)

//КлючНайти -- ищет указанный ключ в свойм списке
func (сам *тПулКлючСлова) КлючНайти(пКлюч мТип.UWordKey) bool {
	for _, слово := range сам.пулСлова {
		if слово.КлючНайти(пКлюч) {
			return true
		}
	}
	return false
}
