package dto

import "github.com/uptrace/bun"

type Restaurant struct {
	bun.BaseModel
	ID       string
	Name     string
	Location string
	AdminID  string
	Staff    []*Staff
	Menu     *Menu
	Tables   []*Table
}
