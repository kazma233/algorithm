package array

/**
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

func plusOne(digits []int) []int {
	length := len(digits) - 1

	for ; length >= 0; length-- {
		if digits[length] == 9 { // 直接制空成0
			digits[length] = 0
		} else { // 加1
			digits[length]++
			return digits
		}
	}
	if digits[0] == 0 { // 全部都是9
		return append([]int{1}, digits...)
	}
	return digits
}
