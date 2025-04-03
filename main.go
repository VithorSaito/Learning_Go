package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

type Storage struct {
	Name string
	Amount int
}

func main() {

	e := echo.New()

	incrementStorage()

	e.GET("/", getStorage)

	e.POST("/create", createItens)

	e.DELETE("/delete/:Name", deleteItens)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "3000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort ))
}

var storage []Storage

func incrementStorage() {
	storage = append(storage, Storage{Name: "beans", Amount: 2})
	storage = append(storage, Storage{Name: "rice", Amount: 10})
	storage = append(storage, Storage{Name: "juice", Amount: 4})
}

func getStorage(c echo.Context) error {

	return c.JSON(200, storage)
}

func createItens( c echo.Context) error {

	itens := new(Storage)
	if err := c.Bind(itens); err != nil {
		return err
	}

	storage = append(storage, *itens)

	return c.JSON(200, storage)
}

func deleteItens (c echo.Context) error {

	name := c.Param("Name")

	for i, item := range storage {
		if item.Name == name {
			storage = append(storage[:i], storage[i+1:]...)
			return c.JSON(200, map[string]string{"message": "Item excluído com sucesso"})
		}
	}
	return c.JSON(200, map[string]string{})
}
