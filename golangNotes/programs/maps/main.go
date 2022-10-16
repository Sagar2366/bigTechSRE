package main

import "fmt"

func main() {

	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Key : %v, Value: %v\n", key, value)
	}
}
