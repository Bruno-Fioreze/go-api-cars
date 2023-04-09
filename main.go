// package main

// import (
// 	"github.com/labstack/echo"
// )

// type Car struct {
// 	Name string
// 	Age  int
// }

// var cars []Car

// func create_cars() []Car {

// 	cars = append(cars, Car{Name: "teste 1", Age: 18})
// 	cars = append(cars, Car{Name: "teste 2", Age: 18})
// 	cars = append(cars, Car{Name: "teste 3", Age: 18})
// 	return cars
// }
// func main() {
// 	cars = create_cars()
// 	e := echo.New()
// 	e.GET("/cars", getCars)
// 	e.POST("/cars", createCar)
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// func getCars(c echo.Context) error {
// 	c.JSON(200, cars)
// 	return nil
// }

// func createCar(c echo.Context) error {
// 	var car Car
// 	if err := c.Bind(&car); err != nil {
// 		return err
// 	}
// 	cars = append(cars, car)
// 	c.JSON(200, cars)
// 	return nil
// }
