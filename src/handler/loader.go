package handler

import (
	"fmt"
	"models"
	"net/http"
	"strconv"
	"strings"
	"util"

	"github.com/labstack/echo"
)

func LoadBookFromCSV(c echo.Context) error {
	dataList := util.LoadFile("data/books.csv")
	count := dataList.Len()
	if count > 0 {
		for e := dataList.Front(); e != nil; e = e.Next() {
			book := models.Book{}
			record := e.Value.([]string)
			book.SetBook(record[0], record[1], record[2], record[3], record[4])
			//book.Print()
			book.Save()
		}
	}

	returnValue := fmt.Sprintf("Total:%d Inserted", count)
	return c.String(http.StatusOK, returnValue)

}

func LoadReadHistoryFromCSV(c echo.Context) error {
	dataList := util.LoadFile("data/readbook.csv")
	count := dataList.Len()
	if count > 0 {
		for e := dataList.Front(); e != nil; e = e.Next() {
			record := e.Value.([]string)

			year := strings.Trim(record[0], " ")
			bookName1 := strings.Trim(record[1], " ")
			bookName2 := strings.Trim(record[2], " ")

			if year == "년" || len(year) == 0 {
				continue
			}
			if len(bookName1) > 0 {
				bookOp := models.BookOp{}
				bookId := strconv.FormatInt(bookOp.FindBookIdByBookName(bookName1), 10)
				bookOp.ReadCSVBookOp(bookId, "이대현", year)

				//fmt.Println("bookName:" + bookName1 + ", " + bookOp.ToString())
				bookOp.Save()
			}
			if len(bookName2) > 0 {
				bookOp := models.BookOp{}
				bookId := strconv.FormatInt(bookOp.FindBookIdByBookName(bookName2), 10)
				bookOp.ReadCSVBookOp(bookId, "이문선", year)

				//fmt.Println("bookName:" + bookName2 + ", " + bookOp.ToString())
				bookOp.Save()
			}
		}
	}

	returnValue := fmt.Sprintf("LoadReadHistoryFromCSV Total:%d Inserted", count)
	return c.String(http.StatusOK, returnValue)

}
