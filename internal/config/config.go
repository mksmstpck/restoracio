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
	CacheExpire   time.Duration
	CachePurge    time.Duration
	TokenSecret []byte
	TokenExp    time.Duration
}

func NewConfig() Config {
	tokenExp, err := strconv.Atoi(os.Getenv("TOKEN_EXP"))
	if err != nil {
		log.Fatal(err)
	}


	cacheExp, err := strconv.Atoi(os.Getenv("CACHE_EXP"))
	if err != nil {
		log.Fatal(err)
	}

	cachePurge, err := strconv.Atoi(os.Getenv("CACHE_PURGE"))
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		GinUrl:        os.Getenv("GIN_URL"),
		CockDNS:       os.Getenv("COCK_DNS"),
		CacheExpire:   time.Duration(cacheExp),
		CachePurge:    time.Duration(cachePurge),
		TokenSecret:   []byte(os.Getenv("TOKEN_SECRET")),
		TokenExp:      time.Duration(tokenExp),
	}
}
