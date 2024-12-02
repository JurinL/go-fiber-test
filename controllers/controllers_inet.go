package controllers

import (
	"log"
	m "go-fiber-test/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"strconv"
)

func HelloTest(c *fiber.Ctx) error { 
	return c.SendString("Hello, World!")
}

//Exercise 5.1
func Factorial(c *fiber.Ctx) error {
	num, err := strconv.Atoi(c.Params("num"))
	result := 1
	if err != nil {
        return err
    }
	for i := 1; i <= num; i++ {
		result *= i
	}
	return c.SendString(c.Params("num") + "! = " + strconv.Itoa(result))
}

func HelloTestV2(c *fiber.Ctx) error { 
	return c.SendString("Hello, World! v2")
}

func BodyParserTest(c *fiber.Ctx) error{
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name)
	log.Println(p.Pass)
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error{
	str := "hello ==> " + c.Params("name")
	return c.JSON(str)	
}

func QueryTest(c *fiber.Ctx) error{
	a := c.Query("search")
	str := "my search is " + a
	return c.JSON(str)
}
// Exercise 5.2
func QueryParams(c *fiber.Ctx) error{
	a := c.Query("tax_id")
	str := ""
	for _, char := range a {
		str += strconv.Itoa(int(char)) + " "
	}
	return c.SendString(str)
}

func ValidTest(c *fiber.Ctx) error {
	
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
		
	}
	return c.JSON(user)
}