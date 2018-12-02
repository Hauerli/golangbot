package main

import (
	"fmt"
)

func checkType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("Type is a string with value ", i.(string))

	case int:
		fmt.Println("Type is an int with value ", i.(int))
	default:
		fmt.Println("Type not known")
	}
}

func main() {

	checkType(67)
	checkType("Hello World")
	checkType(23.2)

}
