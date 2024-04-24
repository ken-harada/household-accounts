package routes

import (
	"fmt"

	"github.com/ken-harada/household-accounts/handlers"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	fmt.Println("Route!")
	e.GET("/", handlers.Index)
	e.GET("/login", handlers.Login)
}
