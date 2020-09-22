package s003

// 给定字符串 S1,S2, 判断字符串 S1 中的字符是否都在字符串 S2 中(S1 长度 N，S2长度 M)
func HasContains(str, substr string) bool {
	slen1 := len(str)
	slen2 := len(substr)
	if slen2 > slen1 {
		return false
	}
	for i := 0; i < slen1; i++ {
		tmp := str[i : i+slen2]
		if len(tmp) < slen2 {
			return false
		}
		if tmp == substr {
			return true
		}
	}
	return false
}
