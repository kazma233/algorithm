package array

import (
	"testing"
)

func Test_containsDuplicate(t *testing.T) {

	arr1 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{1, 2, 3, 1}

	t.Logf("result1: %v", containsDuplicate(arr1))
	t.Logf("result2: %v", containsDuplicate(arr2))
	t.Logf("result3: %v", containsDuplicate(arr3))
}
