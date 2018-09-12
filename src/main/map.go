package main

import "fmt"

func main() {
	test_map()
	test_map2()
	test_map3()
}

func test_map()  {

	map1 := make(map[string]string)

	map1["1"] = "Saprk"
	map1["2"] = "Hadoop"
	map1["3"] = "Storm"

	fmt.Printf("%T,%v,%d\n",map1,map1,len(map1))

}

func test_map2()  {

	map2 := make(map[string]string,32)

	map2["1"] = "Saprk"
	map2["2"] = "Hadoop"
	map2["3"] = "Storm"

	fmt.Printf("%T,%v,%d\n",map2,map2,len(map2))

}

func test_map3()  {

	map3 := map[int]int{
		1:111,
		2:222,
		3:333,
	}

	fmt.Printf("%T,%v,%d\n",map3,map3,len(map3))

}
