package main

import "fmt"

type myFun func(int) int

func main() {
	closure := test_closure()
	for i:=1;i<=100;i++{
		fmt.Println(closure(i))
	}
}

func test_closure() myFun{
	sum := 0

	return func(number int) int{
		sum += number
		return sum
	}
}
