package main

import "fmt"

// 声明全局变量
var gA int = 100
var gB = 200
//gC := 300 不支持全局变量

func main()  {
	// int默认值 = 0
	var a int
	fmt.Println("a = ",a)

	var b int = 100
	fmt.Println("b = ",b)

	// 初始化可以省略数据变量，自动匹配当前变量的数据类型
	var c = 200
	fmt.Println("c = ",c)

	fmt.Printf("type of c = %T\n",c)

	var bb string = "abcd"
	var cc = "abcd"
	fmt.Printf("bb = %s,type of bb = %T\n",bb,bb)
	fmt.Printf("cc = %s,type of cc = %T\n",cc,cc)

	// 省略var关键字，直接自动匹配 (常用方法)
	e := 100
	fmt.Printf("e = %T\n",e)
	fmt.Println("e = ",e)

	fmt.Println("gA = ",gA," gB = ",gB)
	//fmt.Println("gC = ",gC)

	// 声明多个变量
	var xx,yy int = 100,200
	fmt.Println("xx = ",xx," yy = ",yy)

	var kk,ll = 200,"ll"
	fmt.Println("kk = ",kk," ll = ",ll)

	// 多行多变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ",vv," jj = ",jj)
}
