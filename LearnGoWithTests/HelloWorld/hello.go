package main

import "fmt"

const enHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return enHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Dan"))
}
