package s002

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMiddle(t *testing.T) {
	arr := []int{7, 11, 15, 1, 2, 3}
	ret := FindMiddleNumber(arr)
	require.EqualValues(t, 5, ret)

	arr = []int{7, 11, 15, 1, 2, 3, 6}
	ret = FindMiddleNumber(arr)
	require.EqualValues(t, 6, ret)
}
