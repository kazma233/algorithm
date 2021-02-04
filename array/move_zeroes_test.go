package array

import (
	"testing"
)

func Test_moveZeroes(t *testing.T) {

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums)
	t.Logf("result1: %v", nums)

	nums1 := []int{0, 0, 1}
	moveZeroes1(nums1)
	t.Logf("result1: %v", nums1)
}
