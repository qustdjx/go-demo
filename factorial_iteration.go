package main

import "fmt"

func factorial_iteration(n int) int {
	res := 1;
	var i int
	for  i = 1; i <= n; i++ {
		res = i * res
	}
	return res
}

func main() {
	num := 15
	fmt.Println(factorial_iteration(num))
}
