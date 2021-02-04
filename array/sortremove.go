package array

// 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			num1 := nums[:i-1]
			num2 := nums[i:]
			nums = append(num1, num2...)
			i--
		}
	}
	return len(nums)
}
