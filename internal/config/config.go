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
	AccessSecret  []byte
	RefreshSecret []byte
	AccessExp     time.Duration
	RefreshExp    time.Duration
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
		AccessSecret:  []byte(os.Getenv("ACCESS_SECRET")),
		RefreshSecret: []byte(os.Getenv("REFRESH_SECRET")),
		AccessExp:     time.Duration(accessExp),
		RefreshExp:    time.Duration(refreshExp),
	}
}
