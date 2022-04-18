package main

import (
	"fmt"

	"myproject/calc"

	"github.com/google/uuid"
)

func main() {
	// use constant from internal package
	n := calc.MagicNumber
	fmt.Println(n)

	// use function from internal package
	result := calc.Add(1, n)
	fmt.Println(result)

	// use external packge
	id := uuid.NewString()
	fmt.Println(id)

	// use function from same package (different file)
	sayHello("john")
}
