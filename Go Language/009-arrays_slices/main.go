package main

import (
	"fmt"
)

func main() {
	//option 1
	var carTypes [3]string
	carTypes[0] = "Toyota"
	carTypes[1] = "Ford"
	carTypes[2] = "Nissan"
	fmt.Println(carTypes[1])
	fmt.Println("carTypesSlice lenght", len(carTypes))

	//option 2
	carTypes2 := [3]string{"Toyota", "Ford", "Nissan"}
	fmt.Println(carTypes2[0])

	carTypesSlice := []string{"Toyota", "Ford", "Nissan"}
	fmt.Println(carTypesSlice[2])

	carTypesSlice = append(carTypesSlice, "Tesla")
	fmt.Println("carTypesSlice lenght", len(carTypesSlice))

	//option 3
	carTypesSliceMake := make([]string, 3)
	carTypesSliceMake[0] = "Toyota"
	carTypesSliceMake[1] = "Ford"
	carTypesSliceMake[2] = "Nissan"

	fmt.Println("carTypesSlice2", carTypesSlice)
	carTypesSlice2 := carTypesSlice[2:4]
	fmt.Println("carTypeSlice2 length", len(carTypesSlice2))
	fmt.Println("slice[2:4]", carTypesSlice2)
}
