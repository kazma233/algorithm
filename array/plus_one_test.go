package array

import (
	"testing"
)

func Test_plusOne(t *testing.T) {

	nums1 := []int{1, 2, 3}

	t.Logf("result1: %v", plusOne(nums1))

	nums2 := []int{9}
	t.Logf("result1: %v", plusOne(nums2))

	nums3 := []int{9, 9}
	t.Logf("result1: %v", plusOne(nums3))

	nums4 := []int{1, 9, 9}
	t.Logf("result1: %v", plusOne(nums4))
}
