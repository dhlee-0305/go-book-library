package handler

import (
	"fmt"
	"models"
	"net/http"
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
