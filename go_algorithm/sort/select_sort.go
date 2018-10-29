package main

import "fmt"

func main() {
	var arr = []int{1,23,5,7,87,1,3,8,45,2,56,34,67,44}
	select_sort(arr)
	fmt.Println(arr)
}

func select_sort(arr []int){
	for i:=0;i< len(arr)-1;i++{
		init := i
		for j:=i+1;j<len(arr);j++ {
			if arr[j] < arr[i]{
				init = j
			}
		}
		arr[i],arr[init] = arr[init],arr[i]
	}
}