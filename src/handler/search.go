package handler

import (
	"container/list"
	"fmt"
	"net/http"

	"models"
	"util"

	"github.com/labstack/echo"
)

func Search(c echo.Context) error {
	bookName := c.Param("bookName")
	fmt.Printf("input bookName:%s\n", bookName)
	dataList := util.LoadFile("data/books.csv")
	bookList := list.New()

	returnValue := ""
	count := parseBook(dataList, bookList)
	if count > 0 {
		for e := bookList.Front(); e != nil; e = e.Next() {
			book := e.Value.(models.Book)
			if book.Equals(bookName) {
				//book.Print()
				returnValue = book.ToString()
				break
			}

		}
	}
	fmt.Println(returnValue)

	return c.String(http.StatusOK, returnValue)
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
