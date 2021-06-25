package main

import "fmt"

// 声明数据类型,是int的别名
type  myint int

// 定义一个结构体
type Book struct {
	title string
	author string
}

func printBook(book Book){

}

// 结构体是值传递
func changeBook(book Book){
	book.title = "java"
}

func changeBookPointer(book *Book){
	book.title = "c++"
}

func main() {
/*	var a myint = 10
	fmt.Println("a = ",a)
	fmt.Printf("type = %T\n",a)*/

	var book1 Book
	book1.title = "Golang"
	book1.author = "zs"

	fmt.Printf("%v\n",book1)

	changeBook(book1)
	fmt.Println("==========")
	fmt.Printf("%v\n",book1)

	fmt.Println("==========")
	changeBookPointer(&book1)
	fmt.Printf("%v\n",book1)
}
