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

	if language == "arabic" {
		return arabicGreetingPrefix + name
	}
	if language == "french" {
		return frenchGreetingPrefix + name
	}
	if language == "spanish" {
		return spanishGreetingPrefix + name
	}
	return englishGreetingPrefix + name
}
