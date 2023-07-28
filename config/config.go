package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	GinUrl         string
	CockDKS        string
	Access_secret  []byte
	Refresh_secret []byte
	Access_exp     time.Duration
	Refresh_exp    time.Duration
}

func NewConfig() Config {
	access_exp, err := strconv.Atoi(os.Getenv("ACCESS_EXP"))
	if err != nil {
		log.Fatal(err)
	}

	refresh_exp, err := strconv.Atoi(os.Getenv("REFRESH_EXP"))
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		GinUrl:         os.Getenv("GIN_URL"),
		CockDKS:        os.Getenv("COCK_DKS"),
		Access_secret:  []byte(os.Getenv("ACCESS_SECRET")),
		Refresh_secret: []byte(os.Getenv("REFRESH_SECRET")),
		Access_exp:     time.Duration(access_exp),
		Refresh_exp:    time.Duration(refresh_exp),
	}
}
