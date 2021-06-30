package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["USA"] = "NewYork"
	cityMap["England"] = "EG"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, " value =", value)
	}

	//删除
	delete(cityMap, "China")

	// 修改
	changeMap(cityMap)

	fmt.Println("===============")
	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, " value =", value)
	}
}

func printMap(cityMap map[string]string) {
	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, " value =", value)
	}
}

func changeMap(cityMap map[string]string) {
	// 地址传递
	cityMap["England"] = "London"
}
