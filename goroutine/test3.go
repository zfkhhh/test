package main

import (
	"fmt"
	"time"
)

/**
 * 有缓冲channel
 */

func main() {
	// 设置有缓冲的channel
	c := make(chan int, 3)

	fmt.Println("len c = ", len(c), " cap c = ", cap(c))

	go func() {
		defer fmt.Println("子go缓冲结束")

		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("子go程正在运行,元素是 =", i, " ： len(c) = ", len(c), " cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("结束")
}
