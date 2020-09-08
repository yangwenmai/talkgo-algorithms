package s001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindTwoSum(t *testing.T) {
	x1, x2 := FindTwoSum([]int{1, 3, 5, 7, 8, 10, 13}, 15)
	require.EqualValues(t, 2, x1)
	require.EqualValues(t, 5, x2)
	x1, x2 = FindTwoSum([]int{1, 2, 3, 7, 8, 10, 12}, 13)
	require.EqualValues(t, 0, x1)
	require.EqualValues(t, 6, x2)
	x1, x2 = FindTwoSum([]int{0, 2, 3, 10, 11, 44, 55, 66, 77}, 55)
	require.EqualValues(t, 0, x1)
	require.EqualValues(t, 6, x2)
}
