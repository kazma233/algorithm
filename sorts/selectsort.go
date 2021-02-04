package sorts

import (
	"algorithm/common"
)

func sortArrSelect(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}

		common.Swap(i, index, arr)
	}

}
