package main

import (
	"fmt"
)

type vehicle struct {
	mpg           int
	numberOfDoors int
}

type car struct {
	vehicle
	color Color
}

type truck struct {
	vehicle
	color Color
}

type manufacture interface {
	getVin() string
}

//custom type
type Color string

const (
	blue  Color = "blue"
	red   Color = "red"
	black Color = "black"
)

func main() {
	aCar := car{}
	fmt.Println(aCar)
	aCar.mpg = 25
	aCar.numberOfDoors = 4
	fmt.Println(aCar)
	aCar.getMpg()

	bCar := car{vehicle{19, 4}, red}
	fmt.Println(bCar)
	bCar.getMpg()

	cCar := car{}
	cCar.color = black
	cCar.mpg = 34
	fmt.Println(cCar)
	cCar.getMpg()

	//call const func
	defaultCar := newCar()
	fmt.Println("Default value car", defaultCar)

	//polymorphic
	aTruck := truck{vehicle{15, 2}, black}
	mans := [...]manufacture{cCar, aTruck}
	for i, _ := range mans {
		fmt.Println("iteration", i, mans[i].getVin())
	}

}

func (v *vehicle) getMpg() {
	println("This vehicle gets", v.mpg, "mpg")
}

//constructor function
func newCar() *car {
	result := car{}
	result.mpg = 20
	result.numberOfDoors = 4
	result.color = black
	return &result
}

//for example polymorphism
func (t truck) getVin() string {
	return "truck VIN"
}

func (c car) getVin() string {
	return "Car VIN"
}
