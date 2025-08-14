package main

import "fmt"

const prefixHelloEnglish = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return prefixHelloEnglish + name + "!"
}

func main() {
	fmt.Printf("%s", Hello("world"))
}
