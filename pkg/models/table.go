package models

type Table struct {
	ID         string `json:"id" bun:",pk"`
	Number     int    `json:"number"`
	Placement  string `json:"placement"`
	MaxPeople  int    `json:"max_people"`
	IfReserved bool   `json:"if_reserved"`
}
