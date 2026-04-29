package main

import "fmt"


const french = "french"
const swahili = "swahili"
const spanish = "spanish"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, " 
const swahiliHelloPrefix = "Niaje, "

func Hello( name string, language string) string{
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case swahili:
		prefix = swahiliHelloPrefix
	}
	
	return prefix + name
}
func main (){
	fmt.Println(Hello("ed", "swahili"))
}