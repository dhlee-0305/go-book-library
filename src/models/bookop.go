package models

import (
	"fmt"
	"strconv"
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

func (b *BookOp) ReadBookOp(bookId int64, userName string, opDate string) {
	b.setBookOp(bookId, userName, "READ", opDate)
}

func (b *BookOp) HoldBookOp(bookId int64, userName string, opDate string) {
	b.setBookOp(bookId, userName, "Hold", opDate)
}

func (b BookOp) Print() {
	fmt.Printf("bookId:%d, userName:%s, opType:%s, opDate:%s\n", b.bookId, b.userName, b.opType, b.opDate)
}

func (b BookOp) ToString() string {
	return "bookId:" + strconv.FormatInt(b.bookId, 10) + ", userName:" + b.userName + ", opType:" + b.opType + ", opDate:" + b.opDate
}
