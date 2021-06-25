package main

import (
	"fmt"
	"time"
)

func main() {
	var a int = 3
	var age = 10
	var name = "zhangsan"
	time.Sleep(1 * time.Second)
	fmt.Print(fmt.Sprint(age,name,a))
}
