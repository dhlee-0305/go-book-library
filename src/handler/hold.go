package handler

import (
	"models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
)

func HoldingBookList(c echo.Context) error {
	//return c.String(http.StatusOK, "team:"+team+", memhber:"+member)
	return c.String(http.StatusOK, "Holding Book List")

}

func HoldedBook(c echo.Context) error {
	bookOp := models.BookOp{}
	bookIdStr := c.FormValue("bookId")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	userName := c.FormValue("userName")
	bookOp.HoldBookOp(bookId, userName)
	bookOp.Print()

	return c.String(http.StatusOK, "Holding Book OK")
}
