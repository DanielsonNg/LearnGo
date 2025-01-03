package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello"
	printMe(printValue)

	numerator := 4
	denominator := 2
	var result, remainder, err = intDivision(numerator, denominator)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("Division result is %v", result)
	// } else {
	// 	fmt.Printf("Division result is %v with remainder of %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("Division result is %v", result)
	default:
		fmt.Printf("Division result is %v with remainder of %v", result, remainder)
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
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zeroo")
		return 0, 0, err
	}

	return numerator / denominator, numerator % denominator, err
}
