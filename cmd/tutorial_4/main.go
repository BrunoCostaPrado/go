package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// var intArr [3]int32
// 	// intArr[1] = 123
// 	// fmt.Println(intArr[0])
// 	// fmt.Println(intArr[1:3])

// 	// fmt.Println(&intArr[0])
// 	// fmt.Println(&intArr[1])
// 	// fmt.Println(&intArr[2])

// 	// var intArr [3]int32 =[3]int32{1,2,3}
// 	// fmt.Println(intArr[0])

// 	// intArr := [3]int32{1, 2, 3}
// 	// fmt.Println(intArr[0])

// 	intArr := [...]int32{1, 2, 3}
// 	fmt.Println(intArr)

// 	var intSlice []int32 = []int32{4, 5, 6}
// 	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
// 	intSlice = append(intSlice, 7)
// 	fmt.Printf("\nThe length is %v with capacity %v \n", len(intSlice), cap(intSlice))

// 	var intSlice2 []int32 = []int32{8, 9}
// 	intSlice = append(intSlice, intSlice2...)
// 	fmt.Println(intSlice)

// 	var intSlice3 []int32 = make([]int32, 3, 8)
// 	fmt.Println(intSlice3)

// 	var myMap map[string]uint8 = make(map[string]uint8)
// 	fmt.Println(myMap)

// 	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
// 	/* "Emily": 31, "Liam": 19, "Ava": 28, "Noah": 42, "Mia": 36, "Isabella": 25, "Ethan": 38, "Charlotte": 22, "Lucas": 48, "Amelia": 29, "Mason": 41, "Harper": 26, "Logan": 35, "Evelyn": 24, "William": 49, "Abigail": 32, "Oliver": 20, "Sophia": 44, "Benjamin": 39, "Lily": 27, "Caleb": 46, "Madison": 30, "Jaxon": 21, "Victoria": 43, "Alexander": 34, "Jessica": 23, "Gabriel": 40, "Samantha": 37, "Michael": 47, "Aubrey": 33, "Christopher": 18, "Riley": 45, "Anthony": 50, "Zoe": 17, "Julian": 16, "Hannah": 15, "Dominic": 14, "Kayla": 13, "Cameron": 12, "Sydney": 11, "Parker": 10, "Julia": 9, "Cooper": 8, "Avery": 7, "Landon": 6, "Peyton": 5, "Gavin": 4, "Sadie": 3, "Jace": 2, "Brooklyn": 1*/
// 	fmt.Println(myMap2["Adam"])

// 	var age, ok = myMap2["Adam"]
// 	if ok {
// 		fmt.Printf("The age is %v \n", age)
// 	} else {
// 		fmt.Println("Invalid Name")
// 	}

// 	for name, age := range myMap2 {
// 		fmt.Printf("The name is %v, Age:%v \n", name, age)
// 	}

// 	for i, v := range intArr {
// 		fmt.Printf("Index: %v , Value: %v \n", i, v)
// 	}
// 	// var i int = 0
// 	//	for i<10{
// 	//		fmt.Println(i)
// 	//		i=i+1
// 	//	}
// 	// for {
// 	// 	if i >= 10 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(i)
// 	// 	i = i + 1
// 	// }

// 	for i := 0; i < -10; i++ {
// 		fmt.Println(i)
// 	}
// 	for i := range 10 {
// 		fmt.Println(i)
// 	}
// }

func main() {
	var n int = 1000000
	// 100000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v \n", timeLoop(testSlice2, n))

}
func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
