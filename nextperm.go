package nextperm

import (
	"fmt"
	"sort"
)

// NextPermutation will computer the next Lexicographical permuation of the give numbers array.
func NextPermutation(nums []int) {

	target := -1
	var swapper int

	for i := len(nums) - 1; i > 0; i-- {

		if nums[i-1] < nums[i] {
			target = i - 1
			break
		}

	}

	if target < 0 {
		sort.Ints(nums)
		return
	}

	for i := len(nums) - 1; i > target; i-- {
		if nums[i] > nums[target] {
			swapper = i
			break
		}
	}

	nums[target], nums[swapper] = nums[swapper], nums[target]

	sort.Ints(nums[target+1:])
	fmt.Printf("nums after permutation %v\n", nums)
}
