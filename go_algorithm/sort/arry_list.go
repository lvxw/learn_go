package main

import "fmt"

type ArrayList struct {
	arr []int
	size int
}

func add(list *ArrayList,el ...int){
	for i:=0;i<len(el);i++{
		list.arr = append(list.arr, el[i])
		list.size += 1
	}
}

func main() {
	list := new(ArrayList)

	add(list,1,2,3,4,5,7,8,9,10)

	fmt.Println(list.arr)
	fmt.Println(list.size)


}
