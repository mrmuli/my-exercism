package main

import "fmt"

// twoSum's pattern matching is limited and not scalable.
func twoSum(nums []int, target int) (int, int) {
	indices := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		indices[0] = nums[i]
		indices[1] = nums[i+1]
		if indices[0] + indices[1] == target {
			return i, i+1
		} else if indices[0] + nums[i + 2] == target {
			return i, i+2
		}
		continue
	}
	return 0, 0
}

// twoSummer's pattern matching is also flawed, goes out of bounds on #25
func twoSummer(nums []int, target int) []int {
	var indices []int
	for i, j := range nums {
		nextIndex := i + 1
		if j + nums[i + 1] == target {
			indices = append(indices, i, nextIndex)
			break
		}
	}
	return indices
}


func main() {
	nums := []int{3,2,3}
	// 3 + 2
	// 3 + 3
	//
	//
	fmt.Println(twoSum(nums, 6))
}
