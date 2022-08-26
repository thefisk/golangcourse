package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	// Any type which has a function called getGreeting (that returns a string)
	// Becomes an honorary member of type bot, meaning any funcs that accept a bot arg
	// Can now call that function too
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// Original Calls
	printGreetingEn(eb)
	printGreetingSp(sb)

	// New Calls
	printGreeting(eb)
	printGreeting(sb)
}

func printGreetingEn(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSp(sb spanishBot){
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// this could be a really complex function, different to Spanish version
	return "Hello there!"
}

// if not using the input variable, it can be omitted below (compare to English version)
func (sb spanishBot) getGreeting() string {	
	return "Hola!"
}


////// REFACTORED printGreeting which takes an interface

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}