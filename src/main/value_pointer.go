package main

import "fmt"

/**
	基本类型、数组、结构体是值类型
	Map、切片是引用类型
 */

func main() {
	arr := [5]int32{111}
	test(arr)
	fmt.Println(arr)

	slice := []int32{111}
	test2(slice)
	fmt.Println(slice)
}

func test(arr [5]int32)  {
	arr[0] = 0
}

func test2(slice []int32)  {
	slice[0] = 0
}
