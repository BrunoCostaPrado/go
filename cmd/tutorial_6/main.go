package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// type owner struct {
// 	name string
// }

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")

	} else {
		fmt.Println("Need more fuel")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Printf("Total miles left: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 100)

	var myEngine1 electricEngine = electricEngine{20, 15}
	fmt.Printf("Total miles left: %v\n", myEngine1.milesLeft())
	canMakeIt(myEngine1, 43)
}
