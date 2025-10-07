package intailizer

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("error in loading env variable")
	}
}
