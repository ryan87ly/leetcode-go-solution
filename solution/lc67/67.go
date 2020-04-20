package lc67

func toInt(b byte) int {
	if b == '0' {
		return 0
	} else {
		return 1
	}
}

func toStr(i int) string {
	if i == 0 {
		return "0"
	} else {
		return "1"
	}
}

func addBinary(a string, b string) string {
	var longer, shorter string
	if len(a) > len(b) {
		longer, shorter = a, b
	} else {
		longer, shorter = b, a
	}
	result := ""
	flag := 0
	for i := 0; i < len(shorter); i++ {
		v1, v2 := toInt(longer[len(longer)-i-1]), toInt(shorter[len(shorter)-i-1])
		sum := v1 + v2 + flag
		result = toStr(sum%2) + result
		flag = sum / 2
	}
	for i := 0; i < len(longer)-len(shorter); i++ {
		v := toInt(longer[len(longer)-len(shorter)-i-1])
		sum := v + flag
		result = toStr(sum%2) + result
		flag = sum / 2
	}
	if flag > 0 {
		result = "1" + result
	}
	return result
}
