package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v,%T \n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of myString is %v ", len(myString))

	var MyRune = 'a'
	fmt.Printf("\nMyRune = %v", MyRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	catStr := strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
