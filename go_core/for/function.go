package main

import "fmt"

func main() {
	re1, re2 := test(3, 6)
	fmt.Println(re1,re2)

	myFun := test2(1, 2)
	re := myFun(2, 4)
	fmt.Println(re)
}


func test(a int, b int) (int,int){
	return a+b, a*b
}


func test2(a int,b int) (func(int,int) int){
	return func(c int,d int) int{
		return a+b*a+b*c+d
	}
}
