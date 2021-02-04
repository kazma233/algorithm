package sorts

// 计数排序
func countSort(nums []int, maxNum int) {
	count := make([]int, maxNum+1)

	for _, num := range nums {
		count[num]++
	}
	var point = 0             // offset
	for i, c := range count { // i 是待排序的值 c是这个值出现了几次
		for k := 0; k < c; k++ {
			nums[point] = i
			point++
		}
	}
}
