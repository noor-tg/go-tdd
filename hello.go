package main

const englishGreetingPrefix = "Hello, "
const arabicGreetingPrefix = "مرحبا, "
const frenchGreetingPrefix = "Bonjour, "
const spanishGreetingPrefix = "Hola, "

func Hello(name string, language string) string {
	if language == "" {
		language = "english"
	}
	if name == "" {
		name = "World"
	}
	prefix := englishGreetingPrefix

	switch language {
	case "arabic":
		prefix = arabicGreetingPrefix
	case "french":
		prefix = frenchGreetingPrefix
	case "spanish":
		prefix = spanishGreetingPrefix
	}

	return prefix + name
}
