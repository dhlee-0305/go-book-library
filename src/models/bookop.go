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
	BookId   int64  `json:"bookId"`
	UserName string `json:"userName"`
	OpType   string `json:"opType"`
	OpDate   string `json:"opDate"`
}

func (b *BookOp) setBookOp(bookId string, userName string, opType string, opDate string) {
	b.BookId, _ = strconv.ParseInt(bookId, 10, 64)
	b.UserName = strings.Trim(userName, " ")
	b.OpType = strings.Trim(opType, " ")
	b.OpDate = strings.Trim(opDate, " ")
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
	fmt.Printf("bookId:%d, userName:%s, opType:%s, opDate:%s\n", b.BookId, b.UserName, b.OpType, b.OpDate)
}

func (b BookOp) ToString() string {
	return "bookId:" + strconv.FormatInt(b.BookId, 10) + ", userName:" + b.UserName + ", opType:" + b.OpType + ", opDate:" + b.OpDate
}

func (b BookOp) Save() (int64, int) {
	var retVal int = http.StatusOK
	var nRow int64 = 0

	dbCon := db.GetConnector()

	fmt.Printf("BookOp.Save[%s]\n", b.ToString())

	insertSql, _ := dbCon.Prepare("INSERT INTO go_book_op(book_id, user_name, op_type, op_date) VALUES(?, ?, ?, ?)")
	result, err := insertSql.Exec(b.BookId, b.UserName, b.OpType, b.OpDate)
	retVal = CheckErr(err)

	if err == nil && retVal == http.StatusOK {
		nRow, _ = result.RowsAffected()
		retVal = CheckResult(nRow, http.StatusOK)
	}

	defer dbCon.Close()

	return nRow, retVal
}
func FindBookOp(bookId int64, userName string) (BookOp, int) {
	var retVal int = http.StatusOK

	dbCon := db.GetConnector()
	selectSql, err :=
		dbCon.Query(`SELECT book_id, user_name, op_type, op_date 
					FROM go_book_op
					WHERE book_id = ?
					AND user_name = ?`, bookId, userName)
	retVal = CheckErr(err)
	if err != nil {
		fmt.Println(err.Error())
	}

	bookOp := BookOp{}
	for selectSql.Next() {
		err = selectSql.Scan(&bookOp.BookId, &bookOp.UserName, &bookOp.OpType, &bookOp.OpDate)
		if err != nil {
			// nothing
		} else {
			retVal = http.StatusNoContent
		}
	}
	defer dbCon.Close()

	return bookOp, retVal
}
