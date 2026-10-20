package models

type Column struct {
	ID          string `json:"id"`
	BoardID     string `json:"boardID"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Cards []string `json:"cards"`
}
