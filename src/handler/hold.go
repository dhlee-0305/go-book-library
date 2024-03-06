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
	// todo
	return c.String(http.StatusOK, "Holding Book List")

}

func DiscardBookReg(c echo.Context) error {
	bookId, _ := strconv.ParseInt(c.FormValue("bookId"), 10, 64)
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.DiscardBookOp(bookId, userName)
	bookOp.Save()

	return c.String(http.StatusOK, bookOp.ToString())
}

func SellBookReg(c echo.Context) error {
	bookId, _ := strconv.ParseInt(c.FormValue("bookId"), 10, 64)
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.SellBookOp(bookId, userName)
	bookOp.Save()

	return c.String(http.StatusOK, bookOp.ToString())
}

func DonateBookReg(c echo.Context) error {
	bookId, _ := strconv.ParseInt(c.FormValue("bookId"), 10, 64)
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.DonateBookOp(bookId, userName)
	bookOp.Save()

	return c.String(http.StatusOK, bookOp.ToString())
}
