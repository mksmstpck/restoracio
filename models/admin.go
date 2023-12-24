package models

type Admin struct {
	ID           string        `bun:"id,pk"`
	Name         string        `bun:"name"`
	Email        string        `bun:"email"`
	PasswordHash string        `bun:"password"`
	Salt         string        `bun:"salt"`
	Restaurant   *Restaurant `bun:"rel:has-one,join:id=admin_id"`
}