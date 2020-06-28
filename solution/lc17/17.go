package lc17

var mapping = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func run(digits string, current string, idx int) []string {
	if len(digits) == idx {
		return []string{current}
	}
	digit := digits[idx]
	values := mapping[digit]
	result := make([]string, 0)
	for _, v := range values {
		result = append(result, run(digits, current+string(v), idx+1)...)
	}
	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return run(digits, "", 0)
}
