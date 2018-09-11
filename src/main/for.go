package main

import "fmt"

func main() {
	test_for()
	test_for2()
}

func test_for()  {
	arr := [5]int32{1,2,3,4,5}

	for index,value := range arr{
		fmt.Println(index,value)
	}
	fmt.Println("----------------------------")

}

func test_for2()  {
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	fmt.Println("----------------------------")
}
