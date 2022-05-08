package main

import (
	"fmt"
)

func main() {
	printPyramid(5)
}

func printPyramid(height int) {
	for lineNumber := 1; lineNumber <= height; lineNumber++ {
		spaces := height - lineNumber
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		stars := lineNumber*2 - 1
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
