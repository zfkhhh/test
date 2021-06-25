package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (this *User) Call(){
	fmt.Println("user is called....")
	fmt.Println("%v\n",this)
}

func main() {
	user := User{1,"zhangsan",23}

	DoFiledAndMthod(user)
}


func DoFiledAndMthod(input interface{}){

	inputType := reflect.TypeOf(input)

	fmt.Println("input type is :",inputType.Name())

	inputValue := reflect.ValueOf(input)

	fmt.Println("input value = ",inputValue)

	// 获取属性字段
	//1. reflect.Type得到NumField,遍历
	//2. 得到每一个Field,数据类型
	//3 .通过field的Interface()得到value

	for i:=0;i <inputType.NumField();i++{
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s : %v = %v\n",field.Name,field.Type,value)
	}

	for i := 0;i <inputType.NumMethod();i++{
		m := inputType.Method(i)
		fmt.Printf("%s : %v\n",m.Name,m.Type)
	}
}