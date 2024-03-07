package models

import (
	"db"
	"fmt"
	"strconv"
	"strings"
	"util"
)

type Book struct {
	bookId    int64
	bookName  string
	editor    string
	publisher string
	buyDate   string
	status    string
}

func (b *Book) SetBook(bookId string, bookName string, editor string, publisher string, buyDate string, status string) {
	b.bookId, _ = strconv.ParseInt(bookId, 10, 64)
	b.bookName = strings.Trim(bookName, " ")
	b.editor = strings.Trim(editor, " ")
	b.publisher = strings.Trim(publisher, " ")
	b.buyDate = util.ConvertDayFormat(buyDate)
	b.status = strings.Trim(status, " ")
}

func (b *Book) Equals(bookName string) bool {
	return b.bookName == bookName
}

func (b Book) Print() {
	fmt.Printf("bookId:%d, bookName:%s, editor:%s, publisher:%s, buyDate:%s, status:%s\n", b.bookId, b.bookName, b.editor, b.publisher, b.buyDate, b.status)
}

func (b Book) ToString() string {
	return "bookId:" + strconv.FormatInt(b.bookId, 10) + ", bookName:" + b.bookName + ", editor:" + b.editor + ", publisher:" + b.publisher + ", buyDate:" + b.buyDate + ", status:" + b.status
}

func (b Book) Save() {
	dbCon := db.GetConnector()
	fmt.Printf("Book.Save[%s]\n", b.ToString())

	insertSql, err := dbCon.Prepare("INSERT INTO go_book(book_id, book_name, editor, publisher, buy_date, status) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertSql.Exec(b.bookId, b.bookName, b.editor, b.publisher, b.buyDate, b.status)
	defer dbCon.Close()
}
