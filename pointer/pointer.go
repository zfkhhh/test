package main

import "fmt"

func change(p *int) {
	*p = 10
}

func main() {
	a := 1
	// ζιδΌ ι
	change(&a)
	fmt.Println("a = ", a)
}
