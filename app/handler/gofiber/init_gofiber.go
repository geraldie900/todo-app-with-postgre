package gofiber

import (
	"fmt"
	"github.com/geraldie900/todo-app/app/handler/gofiber/apiv1"
	"github.com/geraldie900/todo-app/config"
	"github.com/geraldie900/todo-app/config/utils"
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitGofiber() {

	app := fiber.New()
	app.Use(logger.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.AppConfig.AppPort)))
}

func printInfo(c *fiber.Ctx) error {
	// return metadata project
	return c.SendString(fmt.Sprintf("== %s v%s ==", utils.AppName, utils.AppVersion))
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return printInfo(c)
	})

	api := app.Group("/api")

	// Add your other routes here
	// API v1
	apiv1.ApiV1(api)
}
