package main

import "fmt"

func foo1() int {
	fmt.Println("func defer")
	return 0
}
func foo2() int {
	fmt.Println("func return")
	return 1
}

func re() int {
	defer foo1()
	return foo2()
}


func main(){
	re()
	// 写入defer关键字,通过压栈实现，当前函数main函数生命周期结束时，依次出栈
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main :: hello go 1")
	fmt.Println("main :: hello go 2")
}
