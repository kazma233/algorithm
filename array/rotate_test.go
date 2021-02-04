package array

import "testing"

func Test_rotate(t *testing.T) {

	k := 10
	nums := []int{1, 3, 5, 20, 3, 4, 7}
	rotate(nums, k)

	t.Logf("%v, %d, %d", nums, cap(nums), len(nums))
}
