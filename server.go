package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// e.POST("/expenses", createUserHandler)
	// e.GET("/expenses", getUsersHandler)

	log.Fatal(e.Start(":2565"))

	// fmt.Println("Please use server.go for main file")
	// fmt.Println("start at port:", os.Getenv("PORT"))
}
