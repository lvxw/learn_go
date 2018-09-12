package main

import "fmt"

type MyPerson struct {
	name string
	age int
	gender bool
}

type MyStudent struct {
	MyPerson
	schoolName string
}

func main() {
	myStu := MyStudent{
		MyPerson:MyPerson{"zhangsan",12,true},
		schoolName:"北航小学",
	}

	fmt.Printf("%T,%+v\n",myStu,myStu)
	fmt.Printf("%v\n",myStu.name)
	fmt.Printf("%v\n",myStu.age)
	fmt.Printf("%v\n",myStu.gender)
	fmt.Printf("%v\n",myStu.schoolName)
}
