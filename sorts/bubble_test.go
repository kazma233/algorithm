package sorts

import "testing"

func Test_bubbleSort(t *testing.T) {
	data := []int{1, 4, 6, 2, 4, 3, 11, 6, 87, 0}
	bubbleSort(data)
	t.Logf("冒泡排序结果: %v \n", data)
}
