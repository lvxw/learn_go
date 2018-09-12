package main

import "fmt"

func main() {
	test_goto()
}

func test_goto()  {
	count := 0
	i:=0
	LOOP: for i<100 {
		i++
		for j:=2; j<i; j++{
			if i%j == 0{
				goto LOOP
			}
		}
		fmt.Print(i,",")
		count++
	}
	fmt.Println("\n",count)

}
