package main

import (
	"fmt"

	monolith_config "github.com/TesyarRAz/penggerak/internal/app/monolith/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
)

func main() {
	env := config.NewDotEnv()
	log := config.NewLogger(env)
	fiber := config.NewFiber(env)
	validate := config.NewValidator()
	db := config.NewDatabase(env, log)

	app := monolith_config.NewApp(&config.BootstrapConfig{
		Fiber:    fiber,
		DB:       db,
		Log:      log,
		Validate: validate,
		Env:      env,
	})

	config.Bootstrap(app)

	webPort := env["WEB_PORT"]
	if err := fiber.Listen(fmt.Sprint(":", webPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
