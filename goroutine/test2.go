package main

import "fmt"

/**
 * goroutine 协程间通信（无缓冲channel）
 */

func main() {

	c := make(chan int)

	go func(a int, b int) {
		fmt.Println("a = ", a, "b = ", b)
		// 将a+b写到c中
		c <- a + b
	}(10, 20)
	// 从c中读，并赋值给Num
	num := <-c

	fmt.Println("num = ", num)
}
