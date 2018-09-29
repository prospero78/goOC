// модСекции
package пакСекции

import (
	"fmt"
)

type ТуСекции struct {
}

func Новый() (секции *ТуСекции, ош error) {
	секции = &ТуСекции{}
	return секции, ош
}

func main() {
	fmt.Println("Hello World!")
}
