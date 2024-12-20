package controllers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

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

	db.Find(&dogs)
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
} //test

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

	r := m.ResultDogData{
		Count:       len(dogs),
		Data:        dataResults,
		Name:        "golang-test",
		Sum_red:     sum_red,
		Sum_green:   sum_green,
		Sum_pink:    sum_pink,
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

	db.Where("dog_id BETWEEN 50 AND 100").Find(&dogs) //Exercise 7.1
	return c.Status(200).JSON(dogs)
}

func AddEmployee(c *fiber.Ctx) error { //project_2
	db := database.DBConn
	var employee m.Employees
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	validate := validator.New()
	errors := validate.Struct(employee)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	// Check username pattern - only a-z A-Z 0-9 _ -
	nameMatch, _ := regexp.MatchString(`^[a-zA-Z]+$`, employee.Name)
	if !nameMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name must contain only letters",
		})
	}
	// Check username pattern - only a-z A-Z 0-9 _ -
	lastnameMatch, _ := regexp.MatchString(`^[a-zA-Z]+$`, employee.LastName)
	if !lastnameMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Lastname must contain only letters",
		})
	}
	// Check phone pattern - no whitespace allowed only numbers
	phoneMatch, _ := regexp.MatchString(`^\d+$`, employee.Tel)
	if !phoneMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Phone must contain only numbers without spaces",
		})
	}
	//check birthday
	layouts := []string{"02/01/2006", "2006-01-02", "02-01-2006"}
	var parsedDate time.Time
	var err error
	for _, layout := range layouts {
		parsedDate, err = time.Parse(layout, employee.Birthday)
		if err == nil {
			break
		}
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid birthday format. Supported formats: DD/MM/YYYY, YYYY-MM-DD, DD-MM-YYYY",
			"error":   err.Error(),
		})
	}
	
	employee.Birthday = parsedDate.Format("2006-01-02") // Standardize output format
	db.Create(&employee)
	return c.Status(201).JSON(&employee)
}

func GetEmployees(c *fiber.Ctx) error {
	db := database.DBConn
	var employee []m.Employees
	db.Find(&employee) //delelete = null
	return c.Status(200).JSON(employee)
}
func UpdateEmployee(c *fiber.Ctx) error {
	db := database.DBConn
	var employee m.Employees
	id := c.Params("id")
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Where("id = ?", id).Updates(&employee)
	return c.Status(200).JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var employee m.Employees
	result := db.Delete(&employee, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).SendString("successfully delete Employee profile")
}

func GetEmployeesJson(c *fiber.Ctx) error { // Exercise 7.2
	db := database.DBConn
	var employees []m.Employees
	sum_genz := 0
	sum_geny := 0
	sum_genx := 0
	sum_babyboomer := 0
	sum_gigeneration := 0
	db.Find(&employees)
	var dataResults []m.EmployeesRes
	for _, v := range employees {
		typeStr := ""
		if v.Age < 24 {
			typeStr = "GenZ"
			sum_genz += 1
		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "GenY"
			sum_geny += 1
		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "GenX"
			sum_genx += 1
		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "Baby Boomer"
			sum_babyboomer += 1
		} else {
			typeStr = "G.I. Generation"
			sum_gigeneration += 1
		}

		d := m.EmployeesRes{
			EmployeeID: v.EmployeeID,
			Name:       v.Name,
			LastName:   v.LastName,
			Birthday:   v.Birthday,
			Age:        v.Age,
			Email:      v.Email,
			Tel:        v.Tel,
			Type:       typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultEmployeeData{
		Count:        len(employees),
		Data:         dataResults,
		Name:         "golang-test",
		GenZ:         sum_genz,
		GenY:         sum_geny,
		GenX:         sum_genx,
		BabyBoomer:   sum_babyboomer,
		Gigeneration: sum_gigeneration,
	}
	return c.Status(200).JSON(r)
}

// Company CRUD

func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company

	db.Find(&company)
	return c.Status(200).JSON(company)
}

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var company []m.Company

	result := db.Find(&company, "id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&company)
}

func CreateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func DeleteCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).SendString("Company successfully deleted")
}

func SearchEmployee(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var employees []m.Employees

	result := db.Where("employee_id LIKE ? OR name LIKE ? OR last_name LIKE ?",
		"%"+search+"%",
		"%"+search+"%",
		"%"+search+"%").
		Find(&employees)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&employees)
}
