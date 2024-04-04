package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	} 
	if language == spanish {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
	
}

func Sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Hello("world", ""))
}
