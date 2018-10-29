package main

import "fmt"

func main() {
	var arr = []int{1,23,5,7,87,1,3,8,45,2,56,34,67,44}
	bubble_sort(arr)
	fmt.Println(arr)
}

func bubble_sort(arr []int){
	for i:=0;i< len(arr)-1;i++{
		for j:=0;j<len(arr)-i-1;j++ {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}