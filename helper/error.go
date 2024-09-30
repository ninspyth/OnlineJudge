package er

import (
	"log"
)

func HandleError(msg string, err error) {
	if err != nil {
		log.Println(msg+":")
		log.Fatal(err)	
	}
}

