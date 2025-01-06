package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// func main() {
// 	var intSlice = []int{1, 2, 3}
// 	fmt.Println(sumSlice(intSlice))

// 	var float32Slice = []float32{1, 2, 3}
// 	fmt.Println(sumSlice(float32Slice))
// }

// func sumSlice[T int | float32 | float64](slice []T) T {
// 	var sum T
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	return sum
// }

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchase []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchase)
}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
	data, _ := os.ReadFile(filepath)
	var loaded = []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}
