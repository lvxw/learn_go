package main

import "fmt"

var println = fmt.Println
const colorRed = "red"


func main() {
	result := fmt.Sprintf("%s", "ww")
	fmt.Printf(result)

	var ss = 123
	println(ss)
	println(colorRed)

	const (
		c=8
		d
		e=7
		f
	)

	println(c,d,e,f)


}