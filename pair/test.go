package main

import "fmt"

/**
 变量结构：type+value (叫pair)
 */

func main() {
	// pair<(static)type:string,vlaue="abc">
	var a string = "abc"

	// pair<(static)type:string,vlaue="abc">
	var allType interface{}
	allType = a

	value,_ := allType.(string)
	fmt.Println(value)
}
