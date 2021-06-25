package main

import "fmt"

func printArr(myArr [4]int)  {
	for index,value := range myArr{
		fmt.Println("index = ",index," value = ",value)
	}
}

func main()  {
	var myArr1 [10]int

	for i := 0;i < len(myArr1);i++{
		fmt.Println(myArr1[i])
	}

	myArr2 := [10]int{1,2,3,4}

	for index,value := range myArr2{
		fmt.Println("index = ",index," value = ",value)
	}

	myArr3 := [4]int{11,22,33,44}
	// 数组是值传递，且[4]int和[10]int类型不同
	//printArr(myArr1)
	printArr(myArr3)
	
}
