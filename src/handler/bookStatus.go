package handler

import (
	"encoding/json"
	"models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
)

func DiscardBookReg(c echo.Context) error {
	bookId := c.FormValue("bookId")
	userName := c.FormValue("userName")

	intBookId, _ := strconv.ParseInt(bookId, 10, 64)
	checkBookOp, _ := models.FindBookOp(intBookId, userName)
	if checkBookOp.OpType == "DISC" {
		checkBookOp.DiscardBookOp(bookId, userName)
		result := models.OpResult{}
		result.SetResult(checkBookOp, http.StatusConflict, models.DB_CONFLICT)
		resultJson, _ := json.Marshal(result)

		return c.String(http.StatusOK, string(resultJson))
	}

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
	result.SetResult(bookOp, retVal, resultMessage)
	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}

func SellBookReg(c echo.Context) error {
	bookId := c.FormValue("bookId")
	userName := c.FormValue("userName")

	intBookId, _ := strconv.ParseInt(bookId, 10, 64)
	book, _ := models.FindBookByBookId(intBookId)
	if book.Status == "판매" || book.Status == "기부" {
		checkBookOp := models.BookOp{}
		checkBookOp.SellBookOp(bookId, userName)
		result := models.OpResult{}
		result.SetResult(checkBookOp, http.StatusConflict, models.DB_CONFLICT)
		resultJson, _ := json.Marshal(result)

		return c.String(http.StatusOK, string(resultJson))
	}

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
	result.SetResult(bookOp, retVal, resultMessage)
	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}

func DonateBookReg(c echo.Context) error {
	bookId := c.FormValue("bookId")
	userName := c.FormValue("userName")

	intBookId, _ := strconv.ParseInt(bookId, 10, 64)
	book, _ := models.FindBookByBookId(intBookId)
	if book.Status == "판매" || book.Status == "기부" {
		checkBookOp := models.BookOp{}
		checkBookOp.DonateBookOp(bookId, userName)
		result := models.OpResult{}
		result.SetResult(checkBookOp, http.StatusConflict, models.DB_CONFLICT)
		resultJson, _ := json.Marshal(result)

		return c.String(http.StatusOK, string(resultJson))
	}

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
	result.SetResult(bookOp, retVal, resultMessage)
	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}
