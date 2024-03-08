package handler

import (
	"encoding/json"
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
	nRow, retVal := bookOp.Save()

	var resultMessage string
	if nRow != 0 && retVal == http.StatusOK {
		// nothing
	} else {
		resultMessage = models.DB_INSERT_FAIL
	}

	result := models.OpResult{}
	result.SetResult(bookOp.ToString(), retVal, resultMessage)
	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))

}