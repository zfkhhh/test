package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	/*	for {
		//ok为true表示channel已经关闭
		if data,ok := <-c ;ok{
			fmt.Println(data)
		}else {
			break
		}
	}*/
	// 用range来迭代不断操作的channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Finished .... ")

}
