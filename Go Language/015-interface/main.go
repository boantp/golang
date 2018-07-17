package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
}

type policeOfficer struct {
	badgeNumber int
	precinct    string
}

type behaviors interface {
	talk() string
	walk() int
	run() int
}

//to manipulate data in struct we using function thats when we use interface
func main() {
	p := new(person)
	fmt.Println(p.talk)
	fmt.Println(p.walk)

	o := new(policeOfficer)
	fmt.Println(o.talk)
	fmt.Println(o.walk)
	o.badgeNumber = 1000
	fmt.Println(o.badgeNumber)
	fmt.Println(o.run)

	fmt.Println(run(15))
}

//person implementation
func (p *person) talk() string {
	return "Hi there!"
}

func (p *person) walk() int {
	return 10
}

//policeofficer implementation
func (o *policeOfficer) talk() string {
	return "Have a good day!"
}

func (o *policeOfficer) walk() int {
	return 20
}

func (o *policeOfficer) run() int {
	return 100
}

//regular function
func run(s int) int {
	return s
}
