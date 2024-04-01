package hello

const (
	englishGreetingPrefix = "Hello, "
	arabicGreetingPrefix  = "مرحبا, "
	frenchGreetingPrefix  = "Bonjour, "
	spanishGreetingPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if language == "" {
		language = "english"
	}
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {

	switch language {
	case "arabic":
		prefix = arabicGreetingPrefix
	case "french":
		prefix = frenchGreetingPrefix
	case "spanish":
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}
