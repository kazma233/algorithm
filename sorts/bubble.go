package sorts

// 冒泡
func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[k] > nums[i] {
				temp := nums[k]
				nums[k] = nums[i]
				nums[i] = temp
			}
		}
	}
}
