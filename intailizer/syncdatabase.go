package intailizer

import (
	"log"

	"github.com/barunnbhattarai01/consultancy_backend/model"
)

//sync database means creating table if not exists and
// adding missing colums and rows

func Syncdatabase() {
	err := DB.AutoMigrate(&model.User{}, &model.Register{}, &model.InterviewDate{})

	if err != nil {
		log.Println("Error while migarating", err)
	} else {
		log.Println("sucessfully migarted")
	}
}
