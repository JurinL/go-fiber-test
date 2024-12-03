package main

import (
	"fmt"
	"go-fiber-test/database"
	"github.com/gofiber/fiber/v2"
	"go-fiber-test/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	m "go-fiber-test/models"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Dogs{})
	database.DBConn.AutoMigrate(&m.Company{})
	database.DBConn.AutoMigrate(&m.Employee{})
 }
 
 

func main() {
	app := fiber.New() // create new Fiber instance
	initDatabase()
	routes.InetRoutes(app)
	app.Listen(":3000") //port
}