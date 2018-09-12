package main

import "fmt"

func main() {
	sli := []int{1, 2, 2,6,3,9,1,2,56,5,3, 4, 5, 6, 7, 8, 9, 1,4,2,45}
	fmt.Println(sli)
	test_bubbleSort(sli)
	fmt.Println(sli)
}

func test_bubbleSort(sli []int)  {
	count := 0
	flag := true
	for i:=0;i<len(sli)-1;i++{
		for j:=i+1;j<len(sli);j++{
			if sli[j] <= sli[i] {
				sli[j],sli[i] = sli[i],sli[j]
				count++
				flag = false
			}
		}
		if flag {
			break
		}
	}
	fmt.Println(count)
}
