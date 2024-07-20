func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			j := nums[i]
			if j > len(nums) || j < 1 {
				break
			}

			m := j - 1

			if nums[i] == nums[m] {
				break
			}

			nums[i], nums[m] = nums[m], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}