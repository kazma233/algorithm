package sorts

import (
	"algorithm/common"
	"testing"
)

// 0.171s
func Test_sortArr(t *testing.T) {

	arr := common.Random(10, 0, 100)
	insertSort(arr)
	t.Logf("eg1: %v \n", arr)
}

// 0.178s
func Test_sortArr2(t *testing.T) {

	arr := common.Random(10, 0, 100)
	insertSwitchSort(arr)
	t.Logf("eg2: %v \n", arr)
}
