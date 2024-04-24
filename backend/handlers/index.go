package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	fmt.Println("Index!")
	return c.Render(http.StatusOK, "index.html", nil)
}
