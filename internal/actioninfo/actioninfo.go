package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		return
	}

	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("parsing error: ", err)
			continue
		}
	}
	info, err := dp.ActionInfo()
	if err != nil {
		log.Println("error getting information: ", err)
	}
	log.Println(info)
}
