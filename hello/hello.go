package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const bahasa = "Bahasa"
const bahasaHelloPrefix = "Halo, "

// Hello print hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case bahasa:
		prefix = bahasaHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
