package main

import (
	"os"
	"time"

	"github.com/gsabadini/go-clean-architecture/infrastructure"
	"github.com/gsabadini/go-clean-architecture/infrastructure/database"
	"github.com/gsabadini/go-clean-architecture/infrastructure/log"
	"github.com/gsabadini/go-clean-architecture/infrastructure/router"
	"github.com/gsabadini/go-clean-architecture/infrastructure/validation"
)

func main() {
	app := infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceZapLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}
