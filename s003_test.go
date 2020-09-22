package s003

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasContains(t *testing.T) {
	s1, s2 := "123", "12"
	ret := hasContains(s1, s2)
	require.True(t, ret)

	s1, s2 = "123", "1234"
	ret = hasContains(s1, s2)
	require.False(t, ret)

	s1, s2 = "123", "123"
	ret = hasContains(s1, s2)
	require.True(t, ret)
}
