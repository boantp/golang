package main

import (
	"fmt"
)

const toyotaHeadQuarters, fordHeadQuarters = "Toyota Motor Corporation", "Ford Motor Company"

//var newHQ string = "newHQ"

func main() {
	newHQ := "Ford HQ"
	const myConst = "no value needed"

	fmt.Println(toyotaHeadQuarters, newHQ, myConst)
}
