package main

import (
	"collection/lists"
	"fmt"
)

func init()  {

}

func main() {
	li := lists.NewLinkedList()

	li.AddHead(0)
	li.AddHead(1,2)
	li.Println()

	li.AddTail(3,4)
	li.AddTail(5)
	li.Println()

	li.Set(5,-1)
	li.Println()

	li.InsertBefore(1,7,8,9)
	li.Println()

	li.InsertAfter(0,33,34)
	li.Println()

	li.Remove(10)
	li.Println()

	fmt.Println(li.Get(0))
}

