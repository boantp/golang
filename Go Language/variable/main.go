package main

import (
	"fmt"
)

/*
//Option 1
var carTypeDist1 string = "Toyota"
var carTypeDist2 = "Ford"
*/

/*
//Option 2
var carTypeDist1, carTypeDist2 = "Toyota", "Ford"
*/

//Option 3
var (
	carTypeDist1 = "Toyota"
	carTypeDist2 = "Ford"
)

func main() {
	fmt.Println(carTypeDist1, carTypeDist2)
}
