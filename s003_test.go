package s003

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasContains(t *testing.T) {
	s1, s2 := "123", "12"
	ret := HasContains(s1, s2)
	require.True(t, ret)

	s1, s2 = "123", "1234"
	ret = HasContains(s1, s2)
	require.False(t, ret)

	s1, s2 = "123", "123"
	ret = HasContains(s1, s2)
	require.True(t, ret)
}

// goos: darwin
// goarch: amd64
// BenchmarkHasContains-8          175624195                6.78 ns/op
// PASS
// ok      command-line-arguments  2.763s
func BenchmarkHasContains(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasContains("1234", "12")
	}
}
