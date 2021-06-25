package main

import "fmt"

// 结构体首字母大写，其他包也可以访问
type Hero struct {
	// 属性名首字母大写，为公有属性
	Name string
	Ad int
	Level int
}

/* // 当前this是调用该方法对象的一个副本
func (this Hero) Show(){
	fmt.Println("Name = ",this.Name)
}

func (this Hero) GetName() string{
	return this.Name
}


func (this Hero) SetName(newName string){

	this.Name = newName
}*/

func (this *Hero) Show(){
	fmt.Println("Name = ",this.Name)
}

func (this *Hero) GetName() string{
	return this.Name
}

// 当前this是调用该方法对象的一个副本
func (this *Hero) SetName(newName string){

	this.Name = newName
}

func main() {

	hero := Hero{Name: "zhangsan",Ad: 100,Level: 10}

	hero.Show()

	hero.SetName("lisi")

	hero.Show()

}
