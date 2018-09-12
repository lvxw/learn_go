package main

import "fmt"

type NullInterface interface {

}

type Bird struct {
	name string
}

func myPrint(arg interface{})  {
	fmt.Printf("%T,%v\n",arg,arg)
}

func main() {
	var bird NullInterface = Bird{"juju"}
	fmt.Printf("%T,%v\n",bird,bird)


	myPrint(12)
	myPrint(Bird{"jiji"})

}
