package main

import "fmt"
/**
	继承、重写
 */
type Human struct {
	Name string
	Sex string
}

func (this *Human) Eat()  {
	fmt.Println("Human eat....")
}

func (this *Human) Walk()  {
	fmt.Println("Human walk....")
}

type SuperMan struct {
	Human

	level int
}

// 重写父类eat方法
func (this *SuperMan) Eat(){
	fmt.Println("SuperMan eat...")
}

// 子类新方法
func (this *SuperMan) Fly(){
	fmt.Println("SuperMan fly...")
}

func (this *SuperMan) print(){
	fmt.Println("name = ",this.Name)
	fmt.Println("sex = ",this.Sex)
	fmt.Println("level = ",this.level)
}

func main() {
	h := Human{"zhangsan","man"}

	h.Eat()
	h.Walk()

	// 定义子类对象
	//superMan := SuperMan{Human{"lisi","woman"},10}
	var superMan SuperMan
	superMan.Name = "wangwu"
	superMan.level = 30
	superMan.Sex = "man"

	// 调用子类重新的Eat方法
	superMan.Eat()
	superMan.print()
}