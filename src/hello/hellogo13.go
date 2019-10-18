package main

import "fmt"

type Book struct {
	title string 
	author string
	subject string
	id int 
}

func main() {

	fmt.Println(Book{"go ", "www", "nanoxiong", 1})
	fmt.Println(Book{title:"go1", author:"www1", subject:"nanoxiong1", id:2})

	var book1 Book = Book{"go2", "www2", "nanoxiong2", 3}
	var book2 Book = Book{title:"nanoxiong3", author:"go3", subject:"www3", id:4}

	var add int = 100

	fmt.Println(&add)


	fmt.Println(&book1)
	fmt.Println(&book2)


	fmt.Println("book1.subject=" + book1.subject)
	fmt.Println("book2.author=" + book2.author)
}