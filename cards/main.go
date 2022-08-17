package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string { // need to explicitly provide return object type
// 	return "Five of Diamonds"
// }