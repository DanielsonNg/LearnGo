package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	var c = make(chan int, 5)
// 	go process(c)
// 	for i := range c {
// 		fmt.Println(i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func process(c chan int) {
// 	defer close(c)
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}

//		fmt.Println("Exiting process")
//	}
var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var web = []string{"walmart.com", "cotco.com", "whilefoods.com"}
	for i := range web {
		go checkChickenPrices(web[i], chickenChannel)
		go checkTofuPrices(web[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(web string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- web
			break
		}
	}
}

func checkTofuPrices(web string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			c <- web
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nText Sent: Found deal on tofu at %v", website)
	}
}
