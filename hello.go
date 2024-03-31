package main

const englishGreetingPrefix = "Hello, "
const arabicGreetingPrefix = "مرحبا, "

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
	return englishGreetingPrefix + name
}
