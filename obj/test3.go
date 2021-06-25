package main

import "fmt"

/**
多态
 */

// 本质上是指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}


func (this *Cat) Sleep(){
	fmt.Println("Cat is sleep")
}
func (this *Cat) GetColor() string{
	return this.color
}
func (this *Cat) GetType() string{
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep(){
	fmt.Println("Dog is sleep")
}
func (this *Dog) GetColor() string{
	return this.color
}
func (this *Dog) GetType() string{
	return "Dog"
}

func ShowAnimal(animal AnimalIF){
	animal.Sleep()
	fmt.Println("color = ",animal.GetColor())
	fmt.Println("type = ",animal.GetType())
}

func main() {

/*	var animal AnimalIF // 接口的数据类型
	cat := Cat{"Black"}

	animal = &cat
	animal.Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep()*/

	cat := Cat{"Black"}
	dog := Dog{"Yellow"}

	ShowAnimal(&cat)
	ShowAnimal(&dog)

}
