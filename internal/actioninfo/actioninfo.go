package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает активность, которую можно разобрать из строки
// и для которой можно сформировать информацию для вывода.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает каждую строку набора данных и выводит информацию об активности.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
