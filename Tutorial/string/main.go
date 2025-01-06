package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := []rune("résumé")
	indexed := myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nLength of String is %v", len(myString))

	myRune := 'a'
	fmt.Printf("\nRune = %v,%T", myRune, myRune)

	var strSlice = []string{"s", "t", "r", "i", "n", "g"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	catStr := strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
