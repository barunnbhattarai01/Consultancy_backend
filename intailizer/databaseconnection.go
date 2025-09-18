package intailizer

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// package level variables(other can accessed just by intializer.DB)
var DB *sql.DB

func Connection() {
	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatalf("error in db url ")
	}

	//open database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error while connecting the database %v", err)
	}

	//verify  coonection
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connected %v", err)
	}

	fmt.Printf("connected to postgreess sucessfully")

	DB = db
}
