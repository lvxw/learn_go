package main

import (
	"fmt"
	"sort"
)

func main()  {
	var sli = make([]int,0)
	for k :=0;k<100000;k += 100{
		count := 0
		for i:=k;i<k+100;i++{
			if flag := isPrimeNumber(i);flag {
				count++
				fmt.Printf("%04d->%v\t",i,flag)
				if count%10 == 0{
					fmt.Println()
				}
			}
		}
		sli = append(sli,count)
		fmt.Printf("\n素数的个数是：%v\n",count)
	}

	fmt.Println(sli)

	sort.Ints(sli)
	fmt.Println(sli)

}

func isPrimeNumber(number int) bool {
	for i:=2;i<number;i++{
		if number%i == 0{
			return false
		}
	}
	return true
}
