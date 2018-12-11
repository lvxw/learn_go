package main

import "fmt"

func main() {
	sclice := []int{8,7,6,5,4,3,2,1}
	fmt.Println(sclice)
	bubbleSort(sclice)
	fmt.Println(sclice)
}

func bubbleSort(slice []int){
	for i:= len(slice)-1; i>0; i-- {
		count := 0
		for j:=0;j<i;j++{
			if slice[j] > slice[j+1]{
				slice[j],slice[j+1] = slice[j+1],slice[j]
				count ++
			}
		}
		if count == 0{
			break
		}
	}
}