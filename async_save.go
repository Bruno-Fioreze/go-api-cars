package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Name string
	Age  int
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("cars.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Car{})
}

func main() {
	initDB()

	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	var cars []Car
	db.Find(&cars)
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	var car Car
	if err := c.Bind(&car); err != nil {
		return err
	}
	db.Create(&car)
	return c.JSON(200, car)
}
