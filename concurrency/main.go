package main

import (
	"fmt"
	"os"
)

func main() {
	cars := fillCars()
	go showCars(cars, "first goroutine")
	go showCars(cars, "second goroutine")
	showCars(cars, "no goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//time.Sleep(2 * time.Second)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func fillCars() map[string]int {
	cars := make(map[string]int)
	cars["jeep"] = 10
	cars["toyota"] = 20
	cars["honda"] = 30
	cars["ford"] = 40
	cars["nissan"] = 50

	return cars
}

type Cars map[string]int

func showCars(c Cars, m string) {
	for key, i := range c {
		fmt.Fprintf(os.Stdout, "[%[1]v]%[2]v = %[3]v%[4]v\n", m, i, key, c[key])
	}
}
