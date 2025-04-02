package main

import (
	"github.com/labstack/echo/v4"
)

type Storage struct {
	Name string
	Amount int
}

func main() {

	e := echo.New()

	e.GET("/", getStorage)
	e.Logger.Fatal(e.Start(":8000"))
}

var storage []Storage

func incrementStorage() {
	storage = append(storage, Storage{Name: "beans", Amount: 2})
	storage = append(storage, Storage{Name: "rice", Amount: 10})
	storage = append(storage, Storage{Name: "juice", Amount: 4})
}


func getStorage(c echo.Context) error {

	incrementStorage()

	return c.JSON(200, storage)
}