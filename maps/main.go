package main

import "fmt"

func main() {
	// First way to declare a map
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
		"blue": "#0000ff",
	}

	// Second way to declare a map
	// No values will result in its zero value (blank map)
	var colors2 map[string]string

	// Third way - use make function
	colors3 := make(map[string]string)

	// Add extra entry
	colors["white"] = "#ffffff"

	// Delete an entry by key
	delete(colors,"red")

	fmt.Println(colors, colors2, colors3)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}