package main

import (
	"net/http"

	"github.com/effortless-technologies/elt-marketplace-api/clients"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Tyler!\n")
	})

	e.GET("/delivery-quote", clients.GetDeliveryQuote)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
