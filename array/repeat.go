package array

// 给定一个整数数组，判断是否存在重复元素。

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}
