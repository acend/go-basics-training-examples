package main

import (
	"fmt"
)

func main() {
	printPyramid(5)
}

func printPyramid(height int) {
	for i := 1; i <= height; i++ {
		spaces := height - i
		for j := 0; j < spaces; j++ {
			fmt.Printf(" ")
		}

		chars := i*2 - 1
		for j := 0; j < chars; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
