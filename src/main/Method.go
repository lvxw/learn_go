package main

import "fmt"

type Employee struct {
	name string
	currency string
	salary float64
}

func (e Employee)printSalary()  {
	fmt.Printf("%v,%v%v\n",e.name,e.currency,e.salary)
}

func main() {
	employee := Employee{"lisi", "￥", 8997.5}
	employee.printSalary()
}