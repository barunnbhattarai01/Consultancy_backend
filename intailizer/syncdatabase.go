package intailizer

import (
	"log"
)

//sync database means creating table if not exists and
// adding missing colums and rows

func Syncdatabase() {
	//userauth table
	createUsertable := `
	create table if not exists userauth(
	email text primary key,
	password text not null
	)
	`

	_, err := DB.Exec(createUsertable)
	if err != nil {
		log.Fatal("failed to exec the userauth table")
	}

	//student info register
	createregistertable := `
	 create table if not exists studentregister(
	 id serial primary key,
	  name text not null,
	  address text not null,
	  phone integer not null,
	  Age integer not null,
     join_date  Date not null
	 )
	 `

	_, err = DB.Exec(createregistertable)
	if err != nil {
		log.Fatalf("failed to exec the register table %v", err)
	}

	//interview date
	createinterviewdate := `
	create table if not exists interviewregister(
	id serial primary key,
	 name text not null,
	  address text not null,
	  date Date not null,
	  images text not null
	)
	`
	_, err = DB.Exec(createinterviewdate)
	if err != nil {
		log.Fatal("failed to exec the interview table")
	}

	log.Printf("table execute sucessfully")

}
