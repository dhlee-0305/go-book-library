package handler

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
)

func HoldingBookList(c echo.Context) error {
	//return c.String(http.StatusOK, "team:"+team+", memhber:"+member)
	return c.String(http.StatusOK, "Holding Book List")

}

func HoldedBook(c echo.Context) error {
	return c.String(http.StatusOK, "Holded Book")

}