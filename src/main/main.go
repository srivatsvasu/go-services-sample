package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Cat Type
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func helloo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello From web server echo")
}

func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading Request body in addCat: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)

	if err != nil {
		log.Printf("Failed Unmarshalling in addCat: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your cat: %#v ", cat)

	return c.String(http.StatusOK, "We got your cat!!")

}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your cat name is : %s\nand his data ype is %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Please specify string or json",
	})
}

func main() {
	fmt.Println("Hello World!!")
	e := echo.New()

	e.GET("/", helloo)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	e.Start(":8080")
}
