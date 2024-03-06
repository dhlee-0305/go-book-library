package models

import (
	"db"
	"fmt"
	"strconv"
	"time"
)

type BookOp struct {
	bookId   int64
	userName string
	opType   string
	opDate   string
}

func (b *BookOp) setBookOp(bookId int64, userName string, opType string, opDate string) {
	b.bookId = bookId
	b.userName = userName
	b.opType = opType
	b.opDate = opDate
}

func (b *BookOp) ReadBookOp(bookId int64, userName string) {
	b.setBookOp(bookId, userName, "READ", b.nowDay())
}

func (b *BookOp) HoldBookOp(bookId int64, userName string) {
	b.setBookOp(bookId, userName, "Hold", b.nowDay())
}

func (b BookOp) Print() {
	fmt.Printf("bookId:%d, userName:%s, opType:%s, opDate:%s\n", b.bookId, b.userName, b.opType, b.opDate)
}

func (b BookOp) ToString() string {
	return "bookId:" + strconv.FormatInt(b.bookId, 10) + ", userName:" + b.userName + ", opType:" + b.opType + ", opDate:" + b.opDate
}

func (b BookOp) nowDay() string {
	now := time.Now()
	nowDay := now.Format("2006-01-02")

	return nowDay
}

func (b BookOp) Save() {
	dbCon := db.GetConnector()

	insertSql, err := dbCon.Prepare("INSERT INTO go_book_op(book_id, user_name, op_type, op_date) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertSql.Exec(b.bookId, b.userName, b.opDate, b.opDate)
	defer dbCon.Close()
}
