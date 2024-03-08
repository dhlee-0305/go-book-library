package models

import (
	"db"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"util"
)

type Book struct {
	BookId    int64  `json:"bookId"`
	BookName  string `json:"bookName"`
	Editor    string `json:"editor"`
	Publisher string `json:"publisher"`
	BuyDate   string `json:"buyDate"`
	Status    string `json:"status"`
}

func (b *Book) SetBook(bookId string, bookName string, editor string, publisher string, buyDate string, status string) {
	b.BookId, _ = strconv.ParseInt(bookId, 10, 64)
	b.BookName = strings.Trim(bookName, " ")
	b.Editor = strings.Trim(editor, " ")
	b.Publisher = strings.Trim(publisher, " ")
	b.BuyDate = util.ConvertDayFormat(buyDate)
	b.Status = strings.Trim(status, " ")
}

func (b *Book) SetBookByOp(bookOp BookOp) {
	b.BookId = bookOp.bookId
	b.Status = b.getStatusByOpType(bookOp.opType)
}

func (b *Book) Equals(bookName string) bool {
	return b.BookName == bookName
}

func (b Book) Print() {
	fmt.Printf("bookId:%d, bookName:%s, editor:%s, publisher:%s, buyDate:%s, status:%s\n", b.BookId, b.BookName, b.Editor, b.Publisher, b.BuyDate, b.Status)
}

func (b Book) ToString() string {
	return "bookId:" + strconv.FormatInt(b.BookId, 10) + ", bookName:" + b.BookName + ", editor:" + b.Editor + ", publisher:" + b.Publisher + ", buyDate:" + b.BuyDate + ", status:" + b.Status
}

func (b Book) Save() (int64, int) {
	var retVal int = http.StatusOK
	var nRow int64 = 0

	dbCon := db.GetConnector()
	fmt.Printf("Book.Save[%s]\n", b.ToString())

	insertSql, _ := dbCon.Prepare("INSERT INTO go_book(book_id, book_name, editor, publisher, buy_date, status) VALUES(?, ?, ?, ?, ?, ?)")
	result, err := insertSql.Exec(b.BookId, b.BookName, b.Editor, b.Publisher, b.BuyDate, b.Status)
	retVal = CheckErr(err)

	if err == nil && retVal == http.StatusOK {
		nRow, _ := result.RowsAffected()
		retVal = CheckResult(nRow, http.StatusOK)
	}
	defer dbCon.Close()
	return nRow, retVal
}

func (b Book) UpdateStatus() (int64, int) {
	var retVal int = http.StatusOK
	var nRow int64 = 0

	dbCon := db.GetConnector()
	fmt.Printf("Book.UpdateStatus[%s]\n", b.ToString())

	insertSql, _ := dbCon.Prepare("UPDATE go_book SET status=? WHERE book_id=?")
	result, err := insertSql.Exec(b.Status, b.BookId)
	retVal = CheckErr(err)

	if err == nil && retVal == http.StatusOK {
		nRow, _ := result.RowsAffected()
		retVal = CheckResult(nRow, http.StatusOK)
	}

	defer dbCon.Close()
	return nRow, retVal
}

func (b Book) getStatusByOpType(opType string) string {
	status := ""
	if opType == "SELL" {
		status = "판매"
	} else if opType == "DONA" {
		status = "기부"
	} else {
		fmt.Println("getStatusByOpType: 잘못된 opType:" + opType)
		status = opType
	}

	return status
}
