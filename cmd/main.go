package main

import (
	"database/sql"
	"fmt"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/config"
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/internal/handlers"
	"github.com/mksmstpck/restoracio/internal/services"
	"github.com/patrickmn/go-cache"
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
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func main() {
	// config
	config := config.NewConfig()

	// cockroachdb
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.CockDNS)))
	db := bun.NewDB(sqldb, pgdialect.New())
	adminDB := database.NewAdminDatabase(db)
	restDB := database.NewDRatabase(db)

	// cache
	cache := cache.New(config.CacheExpire, config.CachePurge)

	// services
	service := services.NewServices(adminDB, restDB, cache)

	// gin
	router := gin.Default()
	handlers.NewHandlers(router,
		service,
		config.AccessSecret,
		config.RefreshSecret,
		config.AccessExp,
		config.RefreshExp).HandleAll()

	router.Run(config.GinUrl)
}
