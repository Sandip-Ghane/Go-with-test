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
	return GretingPrefix(language) + name
}

// GretingPrefix ...
func GretingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishPrefix
	case "french":
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("sandip", ""))
}
