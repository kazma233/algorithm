package sorts

import "testing"

func Test_countSort(t *testing.T) {
	arr := []int{1, 4, 6, 2, 4, 3, 11, 6, 87, 0}
	countSort(arr, 87)

	t.Logf("计数排序: %v \n", arr)
}
