package main

import "fmt"

func change(p *int) {
	*p = 10
}

func main() {
	a := 1
	// 指针传递
	change(&a)
	fmt.Println("a = ", a)
}
