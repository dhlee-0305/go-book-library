package handler

import (
	"container/list"
	"db"
	"encoding/json"
	"fmt"
	"net/http"

	"models"
	"util"

	"github.com/labstack/echo"
)

func Search(c echo.Context) error {
	bookName := c.Param("bookName")
	book, resultCode := models.FindBookByBookName(bookName)

	var resultMessage string
	if resultCode == http.StatusOK && len(book.BookName) == 0 {
		resultMessage = models.DB_NO_CONTENT
	}
	searchResult := models.SingleBookResult{}
	searchResult.SetResult(book, resultCode, resultMessage)
	resultJson, _ := json.Marshal(searchResult)

	return c.String(http.StatusOK, string(resultJson))
}

func SearchCSV(c echo.Context) error {
	bookName := c.Param("bookName")
	dataList := util.LoadFile("data/books.csv")
	bookList := list.New()

	searchResult := models.Book{}
	var resultMessage string
	var resultCode int
	var notFound bool = true
	count := parseBook(dataList, bookList)
	if count > 0 {
		for e := bookList.Front(); e != nil; e = e.Next() {
			book := e.Value.(models.Book)
			if book.Equals(bookName) {
				searchResult = book
				notFound = false
				break
			}

		}
	}
	if notFound {
		resultCode = http.StatusNoContent
		resultMessage = models.DB_NO_CONTENT
	}

	result := models.SingleBookResult{}
	result.SetResult(searchResult, resultCode, resultMessage)
	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}

func parseBook(dataList *list.List, bookList *list.List) int {
	count := 0
	for e := dataList.Front(); e != nil; e = e.Next() {
		count++

		book := models.Book{}
		record := e.Value.([]string)
		book.SetBook(record[0], record[1], record[2], record[3], record[4], record[5])
		bookList.PushBack(book)
	}
	fmt.Printf("book count:%d\n", count)

	return count
}

func SearchBookToRead(c echo.Context) error {
	var errCode int = http.StatusOK
	userName := c.Param("userName")

	dbCon := db.GetConnector()
	selectSql, err := dbCon.Query(`SELECT a.book_id, a.book_name, a.editor, a.publisher, a.buy_date, a.status 
									FROM go_book a 
									WHERE a.book_id NOT IN ( 
										SELECT book_id 
										FROM go_book_op 
										WHERE user_name = ? 
									) 
									AND a.status NOT IN ('판매', '기부') 
									AND a.book_name NOT IN('어린왕자', '모비딕')`, userName)
	errCode = models.CheckErr(err)

	result := models.MultiBookResult{}

	for selectSql.Next() {
		book := models.Book{}
		err = selectSql.Scan(&book.BookId, &book.BookName, &book.Editor, &book.Publisher, &book.BuyDate, &book.Status)
		result.Data = append(result.Data, book)

		if err == nil && errCode == http.StatusOK {
			errCode = models.CheckResult(book.BookId, http.StatusNoContent)
		}
	}

	if len(result.Data) == 0 {
		result.ResultMessage = models.DB_NO_CONTENT
	}
	result.ResultCode = errCode
	resultJson, _ := json.Marshal(result)

	defer dbCon.Close()

	return c.String(http.StatusOK, string(resultJson))
}
