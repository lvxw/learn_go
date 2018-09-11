package main

import "fmt"

var a int = 12
const b int = 32

/**
局部变量必须被使用，否则编译报错；全局变量可以不被使用
全局、局部常量,不被使用，编译都不会报错
 */
func main() {
	test_var()
	test_const()
	test_iota()
}

func test_var()  {
	const b int = 32
	var a int = 12
	fmt.Println(a)
}

func test_const()  {
	const b int = 32
}

func test_iota()  {
	const (
		a =1
		b = iota
		c = 2
		d = iota
		e
	)

	const kk = iota

	fmt.Printf("%v,%v,%v,%v,%v\n",a,b,c,d,e)
	fmt.Printf("%v",kk)
}
