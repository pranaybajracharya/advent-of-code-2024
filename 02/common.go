package main

func ascSafe(nums []int) bool {
	for i := 0; i < (len(nums) - 1); i++ {
		if nums[i] == nums[i+1] {
			return false
		} else if nums[i] > nums[i+1] {
			return false
		} else if (nums[i+1] - nums[i]) > 3 {
			return false
		}
	}
	return true
}

func descSafe(nums []int) bool {
	for i := 0; i < (len(nums) - 1); i++ {
		if nums[i] == nums[i+1] {
			return false
		} else if nums[i] < nums[i+1] {
			return false
		} else if (nums[i] - nums[i+1]) > 3 {
			return false
		}
	}
	return true
}

func safe(nums []int) bool {
	if nums[0] == nums[1] {
		return false
	}

	if nums[0] > nums[1] {
		return descSafe(nums)
	}

	return ascSafe(nums)
}
