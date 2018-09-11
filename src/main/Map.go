package main

import "fmt"

func main() {
	test_map()
}

func test_map()  {

	map1 := make(map[string]string)

	map1["1"] = "Saprk"
	map1["2"] = "Hadoop"
	map1["3"] = "Storm"

	for k,v := range map1{
		fmt.Printf("%v -> %v\n",k,v)
	}

	fmt.Println(map1["1"])

	fmt.Println(map1)

}
