package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	c "go-fiber-test/controllers" //Exercise 5.3
)

func InetRoutes(app *fiber.App) {
	// Basic auth middleware
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566", // Exercise 5.0
		},
	}))
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.ParamsTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidTest)
	v1.Get("/fact/:num", c.Factorial)
	v1.Post("/register", c.Register)

	v2 := api.Group("/v2")
	v2.Get("/", c.HelloTestV2)

	v3 := api.Group("/v3")
	v3.Post("/jurin", c.AsciiConvert)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	dog.Get("/delete", c.GetDelete)
}
