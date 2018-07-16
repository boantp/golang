package main

import "fmt"

func main() {
	carDealerInventory := make(map[string]int)
	carDealerInventory["A1 Auto"] = 76
	fmt.Println("A1 Auto Inv =", carDealerInventory["A1 Auto"])
	fmt.Println("length of CarDealerInventory =", len(carDealerInventory))

	carDealerInventory["Southwest Auto Mall"] = 112
	fmt.Println("length of CarDealerInventory =", len(carDealerInventory))
	delete(carDealerInventory, "A1 Auto")
	fmt.Println("length of CarDealerInventory after delete =", len(carDealerInventory))
	fmt.Println("A1 Auto =", carDealerInventory["A1 Auro"])
	fmt.Println("Southwest Auto Mall = ", carDealerInventory["Southwest Auto Mall"])

	carDealerTown := make(map[string]string)
	carDealerTown["A1 Auto"] = "Jakarta"
	carDealerTown["Southwest Auto Mall"] = "Bandung"
	fmt.Println("A1 Auto located in ", carDealerTown["A1 Auto"])
}
