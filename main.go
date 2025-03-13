package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {

	calc := n1 + n2

	return calc
}

func main() {

	text := "A soma do numero é:"

	fmt.Println(
	text,
	sum(1, 10),
	)
}