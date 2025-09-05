package main

import "fmt"

func main() {
	// var p *int32 = new(int32)
	// var i int32

	// fmt.Printf("The value of p point to is: %v \n", *p)
	// fmt.Printf("\nthe value of i is: %v \n", i)
	// *p = 10
	// fmt.Printf("\n The new value of p point to is: %v \n", *p)
	// p = &i
	// *p = 1
	// fmt.Printf("The value of p point to is: %v \n", *p)
	// fmt.Printf("\nthe value of i is: %v \n", i)
	// var k int32=2
	// i=k
	// fmt.Printf("\nthe value of i is: %v \n", i)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n The memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is : %v", result)
	fmt.Printf("\nThe value of thing1 is : %v", thing1)

}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
