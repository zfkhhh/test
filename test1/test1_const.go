package main

import "fmt"

const (
	// iota关键字，每行会累加1，第一行iota默认为0 (iota 只能出现在const中)
	BeiJing = iota
	ShangHai
	ShenZhen

	/*
	BeiJing = 0
	ShangHai = 1
	ShenZhen = 2
	*/
)

const (
	a,b = iota + 1,iota + 2 // iota = 0, a=iota + 1= 1,b=iota + 2=2
	c,d					 	// iota = 1, c=iota + 1= 2,d=iota + 2=3
	e,f 					// iota = 2
	g,h = iota * 2,iota * 3 // iota = 3, g= iota * 2=6,h= iota * 3= 9
	i,j						// iota = 4, i= iota * 2= 8,j= iota * 3 = 12
)

func main()  {
	// 常量,不允许被修改
	const length int = 10

	fmt.Println("length = ",length)

	fmt.Println("BeiJin = ",BeiJing)
	fmt.Println("ShangHai = ",ShangHai)
	fmt.Println("ShenZhen = ",ShenZhen)

	fmt.Println("a = ",a,"b = ",b,"c = ",c,"d = ",d,"e = ",e,"f = ",f,"g = ",g,"h = ",h,"i = ",i,"j = ",j)
}
