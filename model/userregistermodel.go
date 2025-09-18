package model

type Register struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address" `
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
	JOIN_DATE string `json:"join_date"`
}
