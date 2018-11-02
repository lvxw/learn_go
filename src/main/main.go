package main

import (
	"collection"
	"fmt"
)

func main() {
	//testArrayList()

	testLinkedList()
}



func testArrayList()  {
	var list collection.Lists = new(collection.ArrayList)
	list.Add(1,2,3,4,5,6)

	list.Update(-1,10)
	list.Update(9,2)

	fmt.Println(list)

	list.Insert(0,111,111,111)
	fmt.Println(list)
}

func testLinkedList()  {
	list := new(collection.LinkedList)
	list.Add(0,1,2,3,4,5,6,7)
	list.Remove(0)
	list.Remove(2)
	list.Remove(4)
	list.Remove(4)
	list.Remove(4)
	list.AddTail(10,12)
	list.AddHead(111,222)
	list.Remove(3)
	list.ForEach()
	list.Update(10,1000)
	list.ForEach()
	list.Remove(7)
	list.Update(-2,10)
	list.Remove(7)
	list.ForEach()
	list.Insert(-2,11,22,33)
	list.Insert(20,44,55,66)
	list.ForEach()
	list.Insert(2,77,88,99)
	list.ForEach()
	list.Remove(0)
	list.ForEach()

}