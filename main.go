package main

import "fmt"

const englishPrefix = "Hello "

// Hello ...
func Hello(name string) string {
	if name == "" {
		return englishPrefix + "world"
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("sandip"))
}
