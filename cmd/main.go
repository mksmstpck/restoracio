package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/cache"
	"github.com/mksmstpck/restoracio/internal/config"
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/internal/handlers"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/internal/services"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func init() {
	log.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.Function), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

//	@title			Restoracio
//	@version		1.0
//	@description	API for restaurant's management

//	@host		restoracio.fly.dev
//	@BasePath	/

//	@securityDefinitions.apikey	JWTAuth
//	@in							header
//	@name						Authorization

func main() {
	// config
	config := config.NewConfig()

	// cockroachdb
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.CockDNS)))
	db := bun.NewDB(sqldb, pgdialect.New())
	database := database.NewDatabase(db)

	// cache
	opt, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		log.Fatal(err)
	}

	rclient := redis.NewClient(opt)

	cache := cache.NewCache(rclient, config.RedisExp)

	// services
	service := services.NewServices(context.TODO(), database, cache)

	// gin
	router := gin.Default()
	handlers.NewHandlers(
		router,
		service,
		config.AccessSecret,
		config.RefreshSecret,
		config.AccessExp,
		config.RefreshExp,
		).HandleAll()

	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Message{Message: "method not allowed"})
	})
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Message{Message: "route not found"})
	})
	
	if err := router.Run(config.GinUrl); err != nil {
		log.Fatal(err)
	}
}
