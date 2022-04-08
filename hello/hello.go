package main

import (
	"fmt"
)

const EnglishHelloPre = "Hello, "
const SpanishHelloPre = "Hola, "
const FrenchHelloPre = "Bonjour, "
const spanish = "Spanish"
const French = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return helloPre(language) + name

}

func helloPre(language string) (pre string) {
	switch language {
	case spanish:
		pre = SpanishHelloPre
	case French:
		pre = FrenchHelloPre
	default:
		pre = EnglishHelloPre
	}
	return
}

func main() {
	fmt.Println(Hello("Frost", "English"))
}
