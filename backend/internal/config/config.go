package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	GinUrl        string
	CockDNS       string
	AccessSecret  []byte
	RefreshSecret []byte
	AccessExp     time.Duration
	RefreshExp    time.Duration
	EmailSender   string
	EmailPassword string
	SMTPHost      string
	SMTPPort      string
	RedisURL      string
	RedisExp 	  time.Duration
	GlobalURL     string
	MongoURL      string
}

func NewConfig() Config {
	accessExp, err := strconv.Atoi(os.Getenv("ACCESS_EXP"))
	if err != nil {
		log.Fatal(err)
	}

	refreshExp, err := strconv.Atoi(os.Getenv("REFRESH_EXP"))
	if err != nil {
		log.Fatal(err)
	}

	redisExp, err := strconv.Atoi(os.Getenv("REDIS_EXP"))
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		GinUrl:        os.Getenv("GIN_URL"),
		CockDNS:       os.Getenv("COCK_DNS"),
		AccessSecret:  []byte(os.Getenv("ACCESS_SECRET")),
		RefreshSecret: []byte(os.Getenv("REFRESH_SECRET")),
		AccessExp:     time.Duration(accessExp),
		RefreshExp:    time.Duration(refreshExp),
		EmailSender:   os.Getenv("EMAIL_SENDER"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
		SMTPHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      os.Getenv("SMTP_PORT"),
		RedisURL:      os.Getenv("REDIS_URL"),
		RedisExp:      time.Duration(redisExp),
		GlobalURL:     os.Getenv("GLOBAL_URL"),
		MongoURL:      os.Getenv("MONGO_URL"),
	}
}
