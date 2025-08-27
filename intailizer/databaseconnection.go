package intailizer

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// package level variables(other can accessed just by intializer.DB)
var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DB_URL")
	//create the connection to databasee and tell gorm to use postgress drivers
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal("Connection error")
	}
	fmt.Printf("connected to supabase sucessfully")

	DB = db
}
