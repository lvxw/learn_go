package main

import (
	"fmt"
)

type NullInterface interface {

}

type Bird struct {
	name string
}

func (b Bird) String() string {
	return "这只鸟的名字叫"+b.name
}

func myPrint(arg interface{})  {
	fmt.Printf("%T,%v\n",arg,arg)
}

func main() {
	var bird NullInterface = Bird{"juju"}
	fmt.Printf("%T,%v\n",bird,bird)

	var bird2 interface{} = Bird{"juju"}
	//b := bird2.(Bird)
	//b.String()
	fmt.Println(bird2)

	myPrint(12)
	myPrint(Bird{"jiji"})

}


