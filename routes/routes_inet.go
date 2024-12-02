package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	c "go-fiber-test/controllers"
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
	v3 := api.Group("/v3")
	v2.Get("/",c.HelloTestV2)
	v1.Get("/",c.HelloTest)
	v1.Post("/",c.BodyParserTest)
	v1.Get("/user/:name",c.ParamsTest)
	v1.Post("/inet",c.QueryTest)
	v1.Post("/valid",c.ValidTest)
	v1.Get("/fact/:num",c.Factorial)
	v3.Post("/jurin",c.QueryParams)
}