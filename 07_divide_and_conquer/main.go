package main

import "fmt"

func merge(leftNumbers []int, rightNumbers []int) []int {
	left, right, current := 0, 0, 0
	result := make([]int, len(leftNumbers)+len(rightNumbers))

	for left < len(leftNumbers) && right < len(rightNumbers) {
		if leftNumbers[left] < rightNumbers[right] {
			result[current] = leftNumbers[left]
			left++
		} else {
			result[current] = rightNumbers[right]
			right++
		}

		current++
	}

	for left < len(leftNumbers) {
		result[current] = leftNumbers[left]
		current++
		left++
	}

	for right < len(rightNumbers) {
		result[current] = rightNumbers[right]
		current++
		right++
	}

	return result
}

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	middle := len(numbers) / 2
	left := mergeSort(numbers[:middle])
	right := mergeSort(numbers[middle:])

	return merge(left, right)
}

func main() {
	fmt.Println(mergeSort([]int{5, 2, 3, 1}))
}
