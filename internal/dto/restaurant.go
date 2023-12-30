package dto

type Restaurant struct {
	ID       string
	Name     string
	Location string
	AdminID  string
	Staff    []*Staff
	Menu     *Menu
	Tables   []*Table
}
