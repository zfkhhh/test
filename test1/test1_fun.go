package main

import "fmt"

func foo(a string){
	fmt.Println("a = ",a)
}

func foo1(a string,b int)int {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	c := 100
	return c
}
// 返回多个返回值，匿名
func foo2(a string,b int) (int,int) {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	c := 100
	return c,c*3
}
// 返回多个返回值，有形参名
func foo3(a string,b int) (r1 int,r2 int) {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	r1 = 100
	r2 = 200
	return
}

func foo4(a string,b int) (r1 , r2 int){
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	r1 = 100
	r2 = 200
	return
}


func main(){
	c := foo1("abc",44)
	fmt.Println("c = ",c)

	d,e := foo2("abc",44)
	fmt.Println("d = ",d,"e = ",e)

	r1, r2 := foo3("hello", 300)
	fmt.Println("r1 = ",r1,"r2 = ",r2)
}
