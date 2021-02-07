package main

import "fmt"

func search(nums []int, target int) int {
	left, right, middle := 0, len(nums)-1, len(nums)/2

	for left <= right {
		switch {
		case nums[middle] == target:
			return middle
		case nums[middle] < target:
			left = middle + 1
		default:
			right = middle - 1
		}

		middle = (right-left)/2 + left
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
