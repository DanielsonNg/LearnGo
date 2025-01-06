package main

import "fmt"

// type gasEngine struct {
// 	mpg     uint8
// 	gallons uint8
// 	owner
// }

// type owner struct {
// 	name string
// }

// func main() {
// 	myEngine := gasEngine{25, 25, owner{"Mi"}}
// 	myEngine.gallons = 2
// 	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

// }

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

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
		fmt.Println("Got it")
	} else {
		fmt.Println("Nuh uh")
	}
}

func main() {
	myEngine := gasEngine{25, 25}
	canMakeIt(myEngine, 50)
	myEngine2 := electricEngine{2, 10}
	canMakeIt(myEngine2, 100)
}
