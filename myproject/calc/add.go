package calc

func Add(a int, b int) int {
	return internalAdd(a, b)
}

func internalAdd(a int, b int) int {
	return a + b
}
