// struct_test.go
package main

import (
	"fmt"
)

type Book struct {
	book_id int
	author  string
	title   string
	subject string
}

func main() {
	var book1 Book
	var book2 Book

	book1.book_id = 1
	book1.author = "lryong1"
	book1.title = "book test1"
	book1.subject = "test1"

	book2.book_id = 2
	book2.author = "lryong2"
	book2.title = "book test2"
	book2.subject = "test2"

	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 subject : %s\n", book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.book_id)
	fmt.Printf("Cal's result is %d\n", book1.cal_pages_nums())
	fmt.Println("==============================")
	fmt.Printf("Book 1 title : %s\n", book2.title)
	fmt.Printf("Book 1 author : %s\n", book2.author)
	fmt.Printf("Book 1 subject : %s\n", book2.subject)
	fmt.Printf("Book 1 book_id : %d\n", book2.book_id)
	fmt.Printf("Cal's result is %d\n", book2.cal_pages_nums())
}

func (a Book) cal_pages_nums() int {
	return a.book_id << 2
}
