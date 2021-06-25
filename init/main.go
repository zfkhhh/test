package main

import (
	// _ 表示匿名别名，导入但并不使用
	_ "init/lib1"
	// mylib2表示别名
	//mylib2 "init/lib2"

	// .表示导入到当前包中，尽量别用，容易歧义
	. "init/lib2"
)

func main()  {
	//lib1.Lib1Test()
	//mylib2.Lib2Test()
	Lib2Test()
}
