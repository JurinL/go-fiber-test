package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	c "go-fiber-test/controllers" //Exercise 5.3
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v3 := api.Group("/v3")
	dog := v1.Group("/dog")
	company := v1.Group("/company")
	
	

	company.Get("/filter/", c.GetCompany)
	company.Get("/", c.GetCompanies)
	company.Post("/", c.CreateCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.DeleteCompany)
	
	employee := v1.Group("/employee")
	employee.Get("/", c.GetEmployees) //project_2
	employee.Get("/json", c.GetEmployeesJson)
	employee.Get("/search", c.SearchEmployee)
	//basic auth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	}))
	employee.Post("/", c.AddEmployee) //project_2
	employee.Put("/:id", c.UpdateEmployee)
	employee.Delete("/:id", c.DeleteEmployee)
	
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.ParamsTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidTest)
	v1.Get("/fact/:num", c.Factorial)
	v1.Post("/register", c.Register)
	

	v2.Get("/", c.HelloTestV2)

	
	v3.Post("/jurin", c.AsciiConvert)

	//CRUD dogs
	
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	dog.Get("/bin", c.GetDelete) 
	dog.Get("/lens", c.GetLens)

	
}
