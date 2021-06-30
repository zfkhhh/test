package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// 切片截取
	s1 := s[:1]
	fmt.Println(s1)
	s2 := s[1:2]
	fmt.Println(s2)
	// 检验，地址传递
	s[1] = 20
	fmt.Println(s2)
}
