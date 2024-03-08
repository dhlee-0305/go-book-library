package handler

import (
	"encoding/json"
	"models"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
)

func HoldingBookList(c echo.Context) error {
	//return c.String(http.StatusOK, "team:"+team+", memhber:"+member)
	// todo
	return c.String(http.StatusOK, "Holding Book List")

}

func DiscardBookReg(c echo.Context) error {
	bookId := c.FormValue("bookId")
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.DiscardBookOp(bookId, userName)
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

func SellBookReg(c echo.Context) error {
	bookId := c.FormValue("bookId")
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.SellBookOp(bookId, userName)
	nRow, retVal := bookOp.Save()

	var resultMessage string
	if nRow != 0 && retVal == http.StatusOK {
		book := models.Book{}
		book.SetBookByOp(bookOp)
		_, retVal = book.UpdateStatus()
	} else {
		resultMessage = models.DB_INSERT_FAIL
	}

	result := models.OpResult{}
	result.SetResult(bookOp.ToString(), retVal, resultMessage)
	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}

func DonateBookReg(c echo.Context) error {
	bookId := c.FormValue("bookId")
	userName := c.FormValue("userName")

	bookOp := models.BookOp{}
	bookOp.DonateBookOp(bookId, userName)
	nRow, retVal := bookOp.Save()

	var resultMessage string
	if nRow != 0 && retVal == http.StatusOK {
		book := models.Book{}
		book.SetBookByOp(bookOp)
		_, retVal = book.UpdateStatus()
	} else {
		resultMessage = models.DB_INSERT_FAIL
	}

	result := models.OpResult{}
	result.SetResult(bookOp.ToString(), retVal, resultMessage)
	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}
