package main

import (
	"github.com/naeemaei/golang-clean-web-api/api"
	"github.com/naeemaei/golang-clean-web-api/config"

	"github.com/naeemaei/golang-clean-web-api/infra/cache"
	database "github.com/naeemaei/golang-clean-web-api/infra/persistence/database"
	"github.com/naeemaei/golang-clean-web-api/infra/persistence/migration"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	// logger.Info(logging.Redis, logging.Startup, "ALGO FALLO DEMO ", nil)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	// Run database migrations
	migration.Up1()

	api.InitServer(cfg)
}
