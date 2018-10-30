package main

import "fmt"

type Animal2 struct{
	name string
	color string
}
func (a *Animal2) sing(){
	fmt.Println("动物在唱歌")
}

type Cat struct{
	Animal2
	age int
}
/*func (c *Cat) sing(){
	fmt.Println("猫在唱歌")
}
*/


func main() {
	cat := new(Cat)
	cat.name="xiaobai"
	cat.color="white"
	cat.age = 22

	cat.sing()
}
