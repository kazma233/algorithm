package array

/**
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 */

// 自己的
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	for i, first := range nums {
		hasEqual := false
		for j, last := range nums {
			if i == j {
				continue
			}
			if first == last {
				hasEqual = true
				break
			}
		}
		if !hasEqual {
			return first
		}
	}
	panic("must have a single num")
}

// 使用map
func singleNumberByMap(nums []int) int {
	dateMap := make(map[int]int)
	for _, v := range nums {
		dateMap[v]++
	}

	for k, v := range dateMap {
		if v == 1 {
			return k
		}
	}
	panic("must have a single num")
}
