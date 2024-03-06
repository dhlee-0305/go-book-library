package handler

import (
	"models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ReadList(c echo.Context) error {
	team := c.FormValue("team")
	member := c.FormValue("member")
	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)

}

func ReadBookReg(c echo.Context) error {
	bookId, _ := strconv.ParseInt(c.FormValue("bookId"), 10, 64)
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.ReadBookOp(bookId, userName)
	bookOp.Print()
	bookOp.Save()

	return c.String(http.StatusOK, bookOp.ToString())

}
