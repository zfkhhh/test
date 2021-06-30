package main

import (
	"fmt"
	"time"
)

func newTask() {

	i := 0
	for {
		i++
		fmt.Println("new GoRoutine : i =", i)
		time.Sleep(1 * time.Second)
	}
}

// 主GoRoutine
func main() {
	//创建一个go程 去执行newTask
	go newTask()

	fmt.Println("main GoRoutine 退出")

	/*	i := 0
		for {
			i++
			fmt.Println("main goroutine : i =",i)
			time.Sleep(1 * time.Second)
		}*/
}
