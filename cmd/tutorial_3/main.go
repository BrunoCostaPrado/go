package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World! \n"
	printMe(printValue)

	var numerator int = 102
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err)
	// } else if remainder == 0 {
	// 	fmt.Printf("The result is %v \n", result)

	// } else {
	// 	fmt.Printf("The result is %v and the remainder is %v \n", result, remainder)
	// }
	// fmt.Printf("The result is %v and the remainder is %v \n", result, remainder)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result is %v \n", result)
	default:
		fmt.Printf("The result is %v and the remainder is %v \n ", result, remainder)
	}
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

}

func printMe(printValue string) {
	fmt.Print(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
