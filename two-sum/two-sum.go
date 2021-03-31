package main

import "fmt"

// twoSum1 pattern matching is flawed, goes out of bounds on #9. Fails on benchmarks.
func twoSum1(nums []int, target int) []int {
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

// second attempt: twoSum2's pattern matching is limited and not scalable. Fails on benchmarks
func twoSum2(nums []int, target int) (int, int) {
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

// third attempt. Learnt something here and solved for the out of bounds on trial 1.
// I fully named vars to keep a good track
// Apparently >>
// Runtime: 8 ms, faster than 9.97% of Go online submissions for Two Sum.
// Memory Usage: 3.2 MB, less than 99.79% of Go online submissions for Two Sum.
// These are stats from leetCode
func twoSum(nums []int, target int) []int {
	var indices []int
	for index, curVal := range nums {
		for nextIndex := index + 1; nextIndex < len(nums); nextIndex++ {
			if curVal + nums[nextIndex] == target {
				indices = append(indices, index, nextIndex)
			}
		}
	}
	return indices
}

func main() {
	nums := []int{3,2,3}
	fmt.Println(twoSum(nums, 6))
}
