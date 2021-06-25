package main

import "fmt"

/**
interface{}是万能类型
 */

func myFunc(arg interface{}){
	fmt.Println("myFunc is called....")
	fmt.Println(arg)

	//interface区分不同数据类型
	// 类型断言 机制
	value,ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string...")
	}else {
		fmt.Println("arg is string,value = ",value)
		fmt.Printf("value tyep is %T\n",value)
	}

}

type Book struct {
	auth string
}


func main() {

	book := Book{"go"}
	myFunc(book)
	myFunc("100")
	myFunc(3.14)
	myFunc("hello")
}
