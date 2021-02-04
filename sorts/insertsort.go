package sorts

import "algorithm/common"

// 插入排序
// 近乎有序的数组，排序很快，因为break

// eg: 1
func insertSort(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			// 当前这个数(也就是待对比的)和前一个数比较
			if arr[j] < arr[j-1] {
				t := j - 1
				common.Swap(j, t, arr)
			} else {
				break
			}
		}
	}

}

// eg: 2 减少数组交换
func insertSwitchSort(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		e := arr[i] // 需要比较和替换的数值
		p := i      // 需要替换的位置
		for j := i; j > 0; j-- {
			t := j - 1
			if e < arr[t] {
				arr[j] = arr[t]
				p = t
			} else {
				break
			}
		}
		arr[p] = e
	}
}
