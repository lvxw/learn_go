package main

import (
	"collection/lists"
	"fmt"
)

func init()  {

}

func testLinkedList(){
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

func testQueue(){
	queue := lists.NewQueue(5)
	queue.EnQueue(0)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.Println()

	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	queue.Println()

	queue.EnQueue(11)
	queue.EnQueue(22)
	queue.EnQueue(23)
	queue.Println()
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	queue.Println()
	fmt.Println(queue.IsEmpty())

	fmt.Println("-----------------------------")
	queue.EnQueue(0)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.Println()
	fmt.Println(queue.IsFull())
	fmt.Println("-----------------------------")
	queue.Clear()
	fmt.Println(queue)
}

func testStack(){
	stack := lists.NewStack(5)
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Println()
	fmt.Println(stack.IsFull())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println("-------------------------")
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Println()
	fmt.Println(stack.IsFull())
	fmt.Println("-------------------------")
	stack.Clear()
	stack.Println()
	fmt.Println(stack.IsEmpty())
}

func main() {
	testQueue()
}

