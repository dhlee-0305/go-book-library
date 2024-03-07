package handler

import (
	"models"
	"net/http"

	"github.com/labstack/echo"
)

func ReadList(c echo.Context) error {
	team := c.FormValue("team")
	member := c.FormValue("member")
	// todo
	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)

}

func ReadBookReg(c echo.Context) error {
	bookId := c.FormValue("bookId")
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.ReadBookOp(bookId, userName)
	bookOp.Save()

	return c.String(http.StatusOK, bookOp.ToString())

}
