package dto

type Admin struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	Password     string
	Salt         string
	Restaurant   *Restaurant
}
