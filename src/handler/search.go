package handler

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"logger"
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
	userName := c.Param("userName")
	logger.Info("SearchBookToRead - userName=" + userName)

	result := models.MultiBookResult{}
	errCode := models.SearchBookByUserName(userName, &result.Data)

	if len(result.Data) == 0 || errCode != http.StatusOK {
		result.ResultMessage = models.DB_NO_CONTENT
	} else {
		errCode = http.StatusOK
	}
	result.ResultCode = errCode

	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}

// @title SearchReadnBook godoc
// @versoin 1.0
// @Description 사용자가 읽은 책 리스트를 페이지 단위로 조회한다.
// @Param userName query string true "사용자 명"
// @Param limit query int true "읽을 개수"
// @Param offset query int true "읽을 순서(offset)"
// @Success 200 {string} string "도서목록 JSON 문자열"
// @Failure 500 {string} string "조회 실패 내역"
// @Router /search/readn [get]
func SearchReadnBook(c echo.Context) error {
	userName := c.QueryParam("userName")
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	offset, _ := strconv.ParseInt(c.QueryParam("offset"), 10, 64)
	logger.Info(fmt.Sprintf("SearchReadnBook - userName=%s, limit=%d, offset=%d", userName, limit, offset))

	result := models.MultiBookResult{}
	errCode := models.SearchReadBook(userName, (int)(limit), (int)(offset), &result.Data)

	if len(result.Data) == 0 || errCode != http.StatusOK {
		result.ResultMessage = models.DB_NO_CONTENT
	} else {
		errCode = http.StatusOK
	}
	result.ResultCode = errCode

	resultJson, _ := json.Marshal(result)

	return c.String(http.StatusOK, string(resultJson))
}
