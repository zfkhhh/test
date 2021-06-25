package main

import "fmt"

func printArr(myArr []int)  {
	// _表示匿名变量
	for _,value := range myArr{
		fmt.Println("value = ",value)
	}
	// 地址传递
	myArr[0] = 100
}

func main()  {
	// 动态数组，切片slice
	myArr1 := []int{1,2,3,4}

	fmt.Printf("myArr1 type is %T\n",myArr1)

	printArr(myArr1)

	fmt.Println("---------------")

	for _,value := range myArr1{
		fmt.Println("value = ",value)
	}

	fmt.Println("len = %d ,slice = %v\n",len(myArr1),myArr1)

	// 声明slice1是一个切片，但并没有给slice1分配空间
	var slice1 []int
	// make函数开辟3个空间,初始值是0
	slice1 = make([]int,3)
	slice1[0] =100
	fmt.Println("len = %d ,slice = %v\n",len(slice1),slice1)

	// 声明slice2同时分配空间
	slice2 := make([]int,3)
	fmt.Println("len = %d ,slice = %v\n",len(slice2),slice2)



}
