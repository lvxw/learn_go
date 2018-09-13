package main

import "fmt"

func main() {
	count := 0
	for i:=0;i<=2018;i++{
		if year, flag := isLeapYear(i);flag {
			count++
			fmt.Printf("%04d->%v\t",year,flag)
			if count%10 == 0{
				fmt.Println()
			}
		}
	}

	fmt.Println("闰年的个数为：",count)
}

func isLeapYear(year int) (int, bool) {
	if year%400==0 || (year%100 != 0 && year%4 == 0){
		return year,true
	}else {
		return year,false
	}
}
