package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Cat Type
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Dog Type
type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Hamster Type
type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func helloo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello From web server echo")
}

func mainadm(c echo.Context) error {
	return c.String(http.StatusOK, "Hello From Admin!!")
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

func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed Processing addDog: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your dog: %#v ", dog)

	return c.String(http.StatusOK, "We got your dog!!")

}

func addHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)

	if err != nil {
		log.Printf("Failed Processing addDog: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your hamster: %#v ", hamster)

	return c.String(http.StatusOK, "We got your hamster!!")

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

	g := e.Group("/admin")

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${status} ${remote_ip} ${method} ${host} ${path} ${latency_human} ]` + "\n",
	}))

	g.GET("/main", mainadm)
	e.GET("/", helloo)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hams", addHamster)
	e.Start(":8080")
}
