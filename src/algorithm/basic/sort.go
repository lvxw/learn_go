package main

import "fmt"

func main() {
	var mySlice = []int{5,3,2,7,3,1,9}
	fmt.Println(mySlice)
	//bubbleSort(mySlice)
	//insertSort(mySlice)
	quickSort(mySlice,0,len(mySlice)-1)
	selectSort(mySlice)
	fmt.Println(mySlice)
}

func bubbleSort(slice []int){
	var sliceLen = len(slice)
	for i:= 0; i<sliceLen; i++ {
		count := 0
		for j:=0;j<sliceLen-i-1;j++{
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

func insertSort(slice []int){
	for i:=0; i<len(slice)-1; i++{
		for j:=i+1; j>0&&slice[j]<slice[j-1] ;j-- {
			slice[j-1],slice[j] = slice[j],slice[j-1]
		}
	}
}


func selectSort(slice []int){
	for i:=0; i< len(slice); i++{
		index := i
		for j:=index+1; j<len(slice); j++{
			if slice[j] < slice[index]{
				index = j
			}
		}
		if index == i{
			continue
		}else{
			slice[index],slice[i] = slice[i],slice[index]
		}
	}

}

func quickSort(slice []int, start int, end int)  {
	if start >= end{
		return
	}else {
		i := start
		j := end

		baseValue := slice[start]

		for i < j {
			for i < j && slice[j] >= baseValue{
				j --
			}
			if i < j{
				slice[i] = slice[j]
				i++
			}

			for i < j && slice[i] < baseValue{
				i++
			}

			if i < j{
				slice[j] = slice[i]
				j --
			}
		}

		slice[i] = baseValue
		quickSort(slice, start, i-1)
		quickSort(slice, i+1, end)
	}


}