package main

import (
	"fmt"

	user_config "github.com/TesyarRAz/penggerak/internal/app/user/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
)

func main() {
	dotenv := config.NewDotEnv()
	log := config.NewLogger(dotenv)
	app := config.NewFiber(dotenv)
	validate := config.NewValidator()
	db := config.NewDatabase(dotenv, log)

	user_config.Bootstrap(&user_config.BootstrapConfig{
		App:      app,
		DB:       db,
		Log:      log,
		Validate: validate,
		Config:   dotenv,
	})

	webPort := dotenv["WEB_PORT"]
	if err := app.Listen(fmt.Sprint(":", webPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
