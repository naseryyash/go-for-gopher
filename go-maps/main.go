package main

import "fmt"

func main() {
	// syntax #1
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
	}

	// syntax #2
	//var colors map[string]string

	// syntax #3
	//colors := make(map[string]string)

	// add
	colors["white"] = "#ffffff"

	fmt.Println(colors)
	// get
	fmt.Println(colors["white"])

	// update
	// colors["white"] = "#fffff9"
	// fmt.Println(colors["white"])

	// delete
	// delete(colors, "white")
	// fmt.Println(colors["white"])
	// fmt.Println(colors)

	// iteration
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color ", color, " is ", hex)
	}
}
