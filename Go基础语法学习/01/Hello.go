package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Echo",""))
}
