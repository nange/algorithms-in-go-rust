package arraylist

func LongestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以s[i]为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以s[i]和s[i+1]为中心的最长回文子串
		s2 := palindrome(s, i, i+1)

		// 取最长回文串
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}

	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 当前是回文，向两边展开，继续判断
		l--
		r++
	}
	// 返回 l, r之间的最长回文
	return s[l+1 : r]
}
