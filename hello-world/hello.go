package main

import (
	"fmt"
)

const prefixHelloEnglish = "Hello, "
const prefixHelloChinese = "Nǐ hǎo, "
const prefixHelloFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "chinese":
		prefix = prefixHelloChinese
	case "french":
		prefix = prefixHelloFrench
	default:
		prefix = prefixHelloEnglish
	}
	return
}

func main() {
	fmt.Printf("%s", Hello("world", ""))
}
