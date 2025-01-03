package main

import "fmt"

func main() {
	var intArr [3]int32
	// intArr[1] = 2
	// fmt.Println(intArr)
	// fmt.Println(intArr[1:3])

	// intArr := [...]int32{1, 2, 3}
	// fmt.Println(intArr)

	// intSlice := []int32{4, 5, 6}
	// fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// intSlice2 := []int32{8, 9}
	// intSlice2 = append(intSlice, intSlice2...)
	// fmt.Println(intSlice2)

	// var intSlice3 []int32 = make([]int32, 3,10)

	// myMap := make(map[string]uint8)
	// fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 32, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	// delete(myMap2, "Adam")
	// var age, ok = myMap2["Adam"]
	// if ok {
	// 	fmt.Printf("The age is %v", age)
	// } else {
	// 	fmt.Println("Invalid Name")
	// }

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	// i := 0
	// for {
	// 	if i > 9 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
