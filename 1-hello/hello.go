package main

import "fmt"

// 避免每次创建字符串实例
const englishHelloPrefix = "Hello, "

// 西班牙语
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

// 法语
const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
