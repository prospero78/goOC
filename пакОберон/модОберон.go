// модОберон
package пакОберон

import (
	"fmt"
)

type ТуОберон struct {
	Компилер byte
	рес      byte
	гуи      byte
	конс     byte
}

func Новый() (оберон *ТуОберон) {
	оберон = new(ТуОберон)
	return оберон
}

func main() {
	fmt.Println("Hello World!")
}
