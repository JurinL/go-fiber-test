package controllers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// Exercise 5.1
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

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name)
	log.Println(p.Pass)
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {
	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is " + a
	return c.JSON(str)
}

// Exercise 5.2
func AsciiConvert(c *fiber.Ctx) error {
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

// Exercise 6
func Register(c *fiber.Ctx) error {
	user := new(m.Register)
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

	// Check email pattern
	emailMatch, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, user.Email)
	if !emailMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email format",
		})
	}

	// Check username pattern - only a-z A-Z 0-9 _ -
	usernameMatch, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, user.Username)
	if !usernameMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username must contain only letters, numbers, underscore or hyphen",
		})
	}

	// Check password pattern - no whitespace allowed any letter
	passwordMatch, _ := regexp.MatchString(`^\S+$`, user.Password)
	if !passwordMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password cannot contain whitespace",
		})
	}

	// Check phone pattern - no whitespace allowed only numbers
	phoneMatch, _ := regexp.MatchString(`^\d+$`, user.Phone)
	if !phoneMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Phone must contain only numbers without spaces",
		})
	}

	// Check website pattern - https + a-z A-Z 0-9 . - + . a-z A-Z min 2
	websiteMatch, _ := regexp.MatchString(`^(https?)?://[a-z0-9.-]+\.[a-z]{2,}$`, user.Website)
	if !websiteMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid website format",
		})
	}
	return c.JSON(user)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error { // Exercise 7.2
	db := database.DBConn
	var dogs []m.Dogs
	sum_red := 0
	sum_green := 0
	sum_pink := 0
	sum_nocolor := 0
	db.Find(&dogs) 
	var dataResults []m.DogsRes
	for _, v := range dogs { 
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sum_red += 1
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sum_green += 1
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sum_pink += 1
		} else {
			typeStr = "no color"
			sum_nocolor += 1
		}

		d := m.DogsRes{
			Name:  v.Name,  
			DogID: v.DogID, 
			Type:  typeStr, 
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Count: len(dogs),
		Data:  dataResults,
		Name:  "golang-test",
		Sum_red: sum_red,
		Sum_green: sum_green,
		Sum_pink: sum_pink,
		Sum_nocolor: sum_nocolor,
	}
	return c.Status(200).JSON(r)
}

func GetDelete(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs) //Exercise 7.0.2
	return c.Status(200).JSON(dogs)
}

func GetLens(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("dog_id BETWEEN 50 AND 100").Find(&dogs) //Exercise 7.1
	return c.Status(200).JSON(dogs)
}