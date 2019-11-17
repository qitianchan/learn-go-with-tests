package hello

const englishHelloPrefix = "Hello, "
const spenishHelloPrefix = "Helo, "
const frenchHelloPrefix = "Bonjour, "
const french = "Franch"
const spenish = "Spenish"


func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := GreetingPrefix(language)
	
	return prefix + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spenish:
		prefix = spenishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
