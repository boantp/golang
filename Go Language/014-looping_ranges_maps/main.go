package main

import (
	"fmt"
)

func main() {
	//looping through maps
	carDealerInventory := make(map[string]int)
	carDealerInventory["A1 Auto"] = 76
	carDealerInventory["Southwest Auto Mall"] = 112

	carDealerInventory2 := make(map[int]string)
	carDealerInventory2[76] = "A1 Auto Mall"
	carDealerInventory2[112] = "Southwest Auto Mall"

	fmt.Println("\nfirst map loop")
	//this doesnt work for maps
	for m := 0; m < len(carDealerInventory2); m++ {
		fmt.Println(carDealerInventory2[m])
	}
	fmt.Println(carDealerInventory2[76])

	fmt.Println("\nsecond correct map loop")
	for key, _ := range carDealerInventory {
		fmt.Println(key, "inventory =", carDealerInventory[key])
	}
}
