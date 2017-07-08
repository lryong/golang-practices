// struct_test2.go

package main

import (
	"fmt"
)

type Book struct {
	author  string
	title   string
	book_id int
	subject string
}

func main() {
	var book1 Book
	book1.author = "lryong1"
	book1.title = "test1"
	book1.subject = "wulala"
	book1.book_id = 1

	var book2 Book
	book2.author = "lryong2"
	book2.title = "test2"
	book2.subject = "wakaka"
	book2.book_id = 2

	gen_inform(&book1)
	gen_inform(&book2)

}

func gen_inform(a *Book) {
	fmt.Printf("Book's title is %s\n", a.title)
	fmt.Printf("Book's author is %s\n", a.author)
	fmt.Printf("Book's book_id is %d\n", a.book_id)
	fmt.Printf("Book's subject is %s\n", a.subject)
	fmt.Printf("Book's information is %d\n", a.try_method())
	fmt.Println("============================")
}

func (a Book) try_method() int {
	return a.book_id << 4
}
