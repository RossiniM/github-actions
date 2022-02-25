package main

import "fmt"

func Sum(n1 int, n2 int) int {
	return n1 + n2
}

func Multiply(n1 int, n2 int) int {
	return n1 * n2
}

func main() {

	fmt.Print(Sum(1, 3))
	fmt.Print("\n")

	fmt.Print(Multiply(1, 3))
}
