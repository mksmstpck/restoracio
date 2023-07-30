package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"

	"github.com/pborman/uuid"
)

func CreateJWT(exp time.Duration, secret []byte, admin_id uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	clanims := token.Claims.(jwt.MapClaims)
	clanims["exp"] = time.Now().Add(exp * time.Hour).Unix()
	clanims["admin_id"] = admin_id.String()
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(token string, secret []byte) (uuid.UUID, error) {
	parsToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Info("invalid token")
			return nil, errors.New("invalid token")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	user_id := parsToken.Claims.(jwt.MapClaims)["user_id"]
	return uuid.Parse(user_id.(string)), nil
}
