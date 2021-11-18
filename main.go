package main

import "fmt"

const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjour "

// Hello ...
func Hello(name, language string) string {
	if name == "" {
		return englishPrefix + "world"
	}
	if language == "spanish" {
		return spanishPrefix + name
	}
	if language == "french" {
		return frenchPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("sandip", ""))
}
