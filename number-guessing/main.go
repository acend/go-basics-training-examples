package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("guess number between 0 and 9")
	number := rand.Intn(10)
	for {
		fmt.Printf("guess number: ")
		line, err := readLine()
		if err != nil {
			fmt.Println("failed to read number:", err)
			os.Exit(1)
		}

		if line == "exit" {
			break
		}

		currentNumber, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("failed to parse number", err)
			continue
		}

		if number == currentNumber {
			fmt.Println("correct")
			break
		}
	}
}

func readLine() (string, error) {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}
