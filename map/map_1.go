package main

import "fmt"

func main() {
	// 1. ====> 声明map
	var myMap1 map[string]string
	// 开辟数据空间
	myMap1 = make(map[string]string,10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// 2.======> 不确定数据空间大小
	myMap2 := make(map[string]string)
	myMap2["1"] = "java"
	myMap2["2"] = "c++"

	fmt.Println(myMap2)


	//3. ======> 直接赋值
	myMap3 := map[string]string{
		"1" : "java",
		"2" : "c++",
		"3" : "python",
	}
	fmt.Println(myMap3)
}
