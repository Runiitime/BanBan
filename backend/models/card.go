package models

type Card struct {
	ID          string `json:"id"`
	ColumnID    string `json:"columnID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
