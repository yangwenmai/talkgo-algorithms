package s001

// 输入一个递增排序的数组和一个数字 S，在数组中查找两个数，使得他们的和正好是 S，
// 如果有多对数字的和等于 S，输出两个数的乘积最小的。
func FindTwoSum(arr []int, s int) (int, int) {
	ret := map[int]int{}
	for i := 0; i < len(arr); i++ {
		ret[arr[i]] = i
	}
	mini, minj := len(arr)-1, len(arr)-1
	for i := 0; i < len(arr); i++ {
		interVal := s - arr[i]
		if j, ok := ret[interVal]; ok {
			if arr[i]*arr[j] < arr[mini]*arr[minj] {
				mini = i
				minj = j
			}
		}
	}
	return mini, minj
}
