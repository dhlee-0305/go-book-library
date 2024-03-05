package models

import (
	"fmt"
	"strconv"
)

type Book struct {
	bookId    int64
	bookName  string
	editor    string
	publisher string
	buyDate   string
}

func (b *Book) SetBook(bookId int64, bookName string, editor string, publisher string, buyDate string) {
	b.bookId = bookId
	b.bookName = bookName
	b.editor = editor
	b.publisher = publisher
	b.buyDate = buyDate
}

func (b *Book) Equals(bookName string) bool {
	return b.bookName == bookName
}

func (b Book) Print() {
	fmt.Printf("bookId:%d, bookName:%s, editor:%s, publisher:%s, buyDate:%s\n", b.bookId, b.bookName, b.editor, b.publisher, b.buyDate)
}

func (b Book) ToString() string {
	return "bookId:" + strconv.FormatInt(b.bookId, 10) + ", bookName:" + b.bookName + ", editor:" + b.editor + ", publisher:" + b.publisher + ", buyDate:" + b.buyDate
}
