package main

import "fmt"

func factorial(n int) int {
	if (n < 2) {
		return 1
	}
	return n * factorial(n - 1)
}

func main() {
	num := 15
	fmt.Println(factorial(num))
}