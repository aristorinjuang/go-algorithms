package main

import "fmt"

func backtrack(result *[][]int, tmp []int, numbers []int) {
	for i := 0; i < len(numbers); i++ {
		start := append([]int{}, numbers[:i]...)
		newNumbers := append(start, numbers[i+1:]...)
		tmp := append(tmp, numbers[i])

		backtrack(result, tmp, newNumbers)
	}

	if len(numbers) == 0 {
		*result = append(*result, tmp)
	}
}

func permutation(numbers []int) [][]int {
	result := [][]int{}

	backtrack(&result, nil, numbers)

	return result
}

func main() {
	fmt.Println(permutation([]int{1, 2, 3}))
}
