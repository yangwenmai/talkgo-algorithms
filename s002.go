package s002

import (
	"fmt"
)

// 一个有序数组，从随即一位截断，把前段放在后边，如 4 5 6 7 1 2 3 求中位数
func FindMiddleNumber(arr []int) int {
	newArr := sort(arr)
	if len(newArr)%2 == 0 {
		// 偶数
		return (newArr[len(newArr)/2-1] + newArr[len(newArr)/2]) / 2
	} else {
		return newArr[len(newArr)/2]
	}
}

func sort(arr []int) []int {
	newArr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			newArr = append(newArr, arr[i+1:]...)
			newArr = append(newArr, arr[:i+1]...)
			break
		}
	}
	return newArr
}
