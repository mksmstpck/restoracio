package utils

import (
	"crypto/sha1"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=,./;[]{}!@#$%^&*()_+"

func PasswordHash(password string) (string, string) {
	papper := string(symbols[rand.Intn(len(symbols))])
	salt := make([]byte, 10)
    for i := range salt {
        salt[i] = symbols[rand.Intn(len(symbols))]
    }

	hashObj := sha1.New()
	hashObj.Write([]byte(password))
	passwordHash := hashObj.Sum(nil)

	hash, err := bcrypt.GenerateFromPassword([]byte(string(passwordHash) + papper + string(salt)), 6)
	if err != nil {
		panic(err)
	}
	return string(hash), string(salt)
}

func CheckPasswordHash(password, hash, salt string) bool {
	hashObj := sha1.New()
	hashObj.Write([]byte(password))
	passwordHash := hashObj.Sum(nil)

	for i := 0; i < len(symbols); i++ {
		err := bcrypt.CompareHashAndPassword(
			[]byte(hash),
			[]byte(string(passwordHash) + string(symbols[i]) + salt),
		)
		if (err == nil) == true {
			return true
		} 
	}
	return false
}
