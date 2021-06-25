package main

import "fmt"

func main() {

	numbers := make([]int,3,5)

	fmt.Printf("len = %d, cap = %d , slice = %v\n",len(numbers),cap(numbers),numbers)

	// 向numbers切片追加元素1
	numbers = append(numbers,1)

	fmt.Printf("len = %d, cap = %d , slice = %v\n",len(numbers),cap(numbers),numbers)

	fmt.Println("===============")
	numbers2 := make([]int,3)
	fmt.Printf("len = %d, cap = %d , slice = %v\n",len(numbers2),cap(numbers2),numbers2)

	numbers2 = append(numbers2,1)
	fmt.Printf("len = %d, cap = %d , slice = %v\n",len(numbers2),cap(numbers2),numbers2)


}
