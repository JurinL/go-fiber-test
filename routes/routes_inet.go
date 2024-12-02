package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"go-fiber-test/controllers"
)

func InetRoutes(app *fiber.App) {
	// Basic auth middleware
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))
	// /api/v1
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v2.Get("/",controllers.HelloTestV2)
	v1.Get("/",controllers.HelloTest)
	v1.Post("/",controllers.BodyParserTest)
	v1.Get("/user/:name",controllers.ParamsTest)
	v1.Post("/inet",controllers.QueryTest)
	v1.Post("/valid",controllers.ValidTest)
}