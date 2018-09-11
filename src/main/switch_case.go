package main

import "fmt"

func main() {
	test_switch_case()
}

func test_switch_case() {

	score := 78
	grade := ""
	comment := ""

	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "{}"
	case score >= 70:
		grade = "C"
		fallthrough
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}

	switch grade {
	case "A":
		comment = "1"
	case "B":
		comment = "2"
	case "C":
		comment = "3"
	case "D":
		comment = "4"
	default:
		comment = "other"
	}

	fmt.Println(grade)
	fmt.Println(comment)

}
