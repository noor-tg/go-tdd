package main

const greetingPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix + name
}
