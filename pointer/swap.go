package main

import "fmt"

func swap(a *int,b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20

	//swap
	swap(&a,&b)

	fmt.Println("a = ",a," b = ",b)

	var p *int
	p = &a

	fmt.Println(&a)
	fmt.Println(p)
	// 二级指针
	var pp **int
	pp = &p
	fmt.Println(pp)
	fmt.Println(&p)
}
