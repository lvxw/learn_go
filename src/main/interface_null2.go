package main

import "fmt"

type A interface {

}


func main() {
	as := []interface{} {1, 2, 34, 4}
	as2 := []A {1, 2, 34, 4}
	fmt.Println(as)
	fmt.Println(as2)

	myMap := map[int]interface{}{
		1: 1,
		2: "2",
		3: []int{1,2,3,4},
	}
	fmt.Println(myMap)
}
