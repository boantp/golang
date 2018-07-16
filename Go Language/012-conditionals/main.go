package main

import (
	"fmt"
)

func main() {
	carLotA := 20
	carLotB := 29

	if carLotB <= carLotA {
		fmt.Println("carLotA is bigger than carLotB")
	} else {
		fmt.Println("condition failed - fall through")
	}

	if carLotA > 15 {
		fmt.Println("carLotA value is bigger than 15")
	} else if carLotA > 9 {
		fmt.Println("carLotA value is bigget than 9")
	} else {
		fmt.Println("Default Case")
	}

	switch carLotA {
	case 15:
		fmt.Println("case carLotA value is 15")
	case 9, 11, 12:
		fmt.Println("case carLotA value is 9 or 11 or 12")
	default:
		fmt.Println("default case--")
	}
}
