package models

import (
	"db"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"util"
)

type BookOp struct {
	bookId   int64
	userName string
	opType   string
	opDate   string
}

func (b *BookOp) setBookOp(bookId string, userName string, opType string, opDate string) {
	b.bookId, _ = strconv.ParseInt(bookId, 10, 64)
	b.userName = strings.Trim(userName, " ")
	b.opType = strings.Trim(opType, " ")
	b.opDate = strings.Trim(opDate, " ")
}

func (b *BookOp) ReadBookOp(bookId string, userName string) {
	b.setBookOp(bookId, userName, "READ", util.NowDay())
}

func (b *BookOp) ReadCSVBookOp(bookId string, userName string, opDate string) {
	b.setBookOp(bookId, userName, "READ", opDate)
}

func (b *BookOp) DiscardBookOp(bookId string, userName string) {
	b.setBookOp(bookId, userName, "DISC", util.NowDay())
}

func (b *BookOp) SellBookOp(bookId string, userName string) {
	b.setBookOp(bookId, userName, "SELL", util.NowDay())
}

func (b *BookOp) DonateBookOp(bookId string, userName string) {
	b.setBookOp(bookId, userName, "DONA", util.NowDay())
}

func (b BookOp) Print() {
	fmt.Printf("bookId:%d, userName:%s, opType:%s, opDate:%s\n", b.bookId, b.userName, b.opType, b.opDate)
}

func (b BookOp) ToString() string {
	return "bookId:" + strconv.FormatInt(b.bookId, 10) + ", userName:" + b.userName + ", opType:" + b.opType + ", opDate:" + b.opDate
}

func (b BookOp) Save() (int64, int) {
	var retVal int = http.StatusOK
	var nRow int64 = 0

	dbCon := db.GetConnector()

	fmt.Printf("BookOp.Save[%s]\n", b.ToString())

	insertSql, _ := dbCon.Prepare("INSERT INTO go_book_op(book_id, user_name, op_type, op_date) VALUES(?, ?, ?, ?)")
	result, err := insertSql.Exec(b.bookId, b.userName, b.opType, b.opDate)
	retVal = CheckErr(err)

	if err == nil && retVal == http.StatusOK {
		nRow, _ := result.RowsAffected()
		retVal = CheckResult(nRow, http.StatusOK)
	}

	defer dbCon.Close()

	return nRow, retVal
}
