package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result, numbers := twoSum(nums, target)
	fmt.Printf("Indices: %v\n", result)
	fmt.Printf("Numbers: %v\n", numbers)
}

func twoSum(nums []int, target int) ([]int, []int) {
	indexMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := indexMap[target-num]; ok {
			return []int{j, i}, []int{nums[j], nums[i]}
		}
		indexMap[num] = i
	}
	return nil, nil
}
