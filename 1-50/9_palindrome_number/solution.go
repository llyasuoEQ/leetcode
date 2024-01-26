package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var (
		pNum = 0
		temp = x
	)
	for temp > 0 {
		pNum = (pNum * 10) + (temp % 10)
		temp = temp / 10
	}
	return pNum == x
}
