package main

import "fmt"

// Supported languages
const (
	greek  = "Greek"
	french = "French"
)

// Language prefixes
const (
	greekPrefix        = "Geia sou "
	englishHelloPrefix = "Hello, "
	frenchPrefix       = "Bonjour, "
)

const defaultName = "World"

func Hello(name, language string) string {

	if name == "" {
		return greetingPrefix(language) + defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case greek:
		return greekPrefix
	case french:
		return frenchPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {

	fmt.Println(Hello("World", "Greek"))
}
