package main

import "fmt"

const enHelloPrefix = "Hello, "

func Hello(name string) string {
	return enHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Dan"))
}
