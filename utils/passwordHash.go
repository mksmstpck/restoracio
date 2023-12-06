package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, string) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=,./;[]{}!@#$%^&*()_+"
	salt := make([]byte, 10)
    for i := range salt {
        salt[i] = letterBytes[rand.Intn(len(letterBytes))]
    }

	hash, err := bcrypt.GenerateFromPassword([]byte(password+string(salt)), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash), string(salt)
}

func CheckPasswordHash(password, hash, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}
