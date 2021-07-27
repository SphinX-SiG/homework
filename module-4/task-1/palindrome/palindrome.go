package palindrome

func IsPalindrome(s string) bool {
	res := []rune(s)

	for i, j := 0, len(res)-1; i < len(res)/2; i, j = i+1, j-1 {
		if res[i] != res[j] {
			return false
		}
	}
	return true
}
