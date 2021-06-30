package main

import (
	"fmt"
	"reflect"
)

/**
反射
*/

func reflectNum(arg interface{}) {
	fmt.Println("type", reflect.TypeOf(arg))
	fmt.Println("value", reflect.ValueOf(arg))

}

func main() {

	var num float64 = 1.234

	reflectNum(num)
}
