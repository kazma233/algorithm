package array

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	var arr = []int{1, 2, 2, 3, 3, 5, 6, 7, 10, 10, 11}
	var last = removeDuplicates(arr)
	t.Logf("result: %v", arr[:last])
}
