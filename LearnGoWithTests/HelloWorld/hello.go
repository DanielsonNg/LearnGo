package main

import (
	"fmt"
)

const (
	enHelloPrefix      = "Hello, "
	spanishHelloPrefix = "Hola, "
	frechHelloPrefix   = "Bonjour, "

	spanish = "Spanish"
	french  = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPRefix(language) + name

}

func greetingPRefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frechHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Dan", ""))
}
