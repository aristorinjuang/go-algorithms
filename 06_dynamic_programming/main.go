package main

import "fmt"

var numbers map[int]int

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciMemoization(n int) int {
	if numbers[n] != 0 {
		return numbers[n]
	}

	if n <= 2 {
		numbers[n] = 1
	} else {
		numbers[n] = fibonacciMemoization(n-1) + fibonacciMemoization(n-2)
	}

	return numbers[n]
}

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	result, less1, less2 := 1, 0, 0

	for i := 2; i <= n; i++ {
		less2 = result
		result += less1
		less1 = less2
	}

	return result
}

func main() {
	numbers = make(map[int]int)
	fmt.Println(fibonacciRecursive(7))
	fmt.Println(fibonacciMemoization(7))
	fmt.Println(fibonacciIterative(7))
}
