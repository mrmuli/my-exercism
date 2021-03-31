package main

import "fmt"

// twoSummer's pattern matching is limited and not scalable.
func twoSummer(nums []int, target int) []int {
	var indy []int
	indices := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		indices[0] = nums[i]
		indices[1] = nums[i+1]
		if indices[0] + indices[1] == target {
			indy = append(indy, i, i+1)
		} else if indices[0] + nums[i + 2] == target {
			indy = append(indy, i, i+2)
		}
		continue
	}
	return indy
}

// twoSum's pattern matching is also flawed, goes out of bounds on #26
func twoSum(nums []int, target int) []int {
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
	fmt.Println(twoSum(nums, 6))
}