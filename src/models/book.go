package models

import "fmt"

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

func (b Book) Print() {
	fmt.Printf("bookId:%d, bookName:%s, editor:%s, publisher:%s, buyDate:%s\n", b.bookId, b.bookName, b.editor, b.publisher, b.buyDate)
}
