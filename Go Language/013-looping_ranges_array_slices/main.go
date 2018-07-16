package main

import (
	"fmt"
)

func main() {
	var carTypes [3]string
	carTypes[0] = "Toyota"
	carTypes[1] = "Ford"
	carTypes[2] = "Nissan"

	//option 1
	i := 0
	for i < len(carTypes) {
		fmt.Println(carTypes[i])
		i++
	}

	//option 2
	fmt.Println("\nfor loop verbose")
	for j := 0; j < len(carTypes); j++ {
		fmt.Println(carTypes[i])
	}

	//option 3
	fmt.Println("\nfor range loop")
	for k, value := range carTypes {
		fmt.Println("k=", k, "&value=", value)
	}

	//option 4
	fmt.Println("\nfor range loop ignore key or value")
	for _, value := range carTypes {
		fmt.Println("k=", "&value=", value)
	}

}
