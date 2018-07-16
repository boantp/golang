package main

import (
	"fmt"
)

type dealerShip struct {
	name string
	city string
}

func main() {
	d1 := dealerShip{"A1 Auto", "Los Angeles"}
	dealerName := createDealerFullName(d1.name, d1.city)
	fmt.Println(dealerName)

	sold, name := calculateInvutil(100, 175, d1)
	fmt.Println(sold, name)
}

func createDealerFullName(s1 string, s2 string) string {
	return s1 + "of" + s2
}

func calculateInvutil(remaining float32, start float32, dealer dealerShip) (percentSold float32, dealerName string) {
	dealerName = dealer.name + "sold"
	percentSold = 1 - remaining/start
	return
}
