package intailizer

import "github.com/barunnbhattarai01/consultancy_backend/model"

//sync database means creating table if not exists and
// adding missing colums and rows

func Syncdatabase() {
	DB.AutoMigrate(&model.User{}, &model.Register{}, &model.InterviewDate{})
}
