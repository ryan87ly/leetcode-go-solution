package lc509

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	prev2 := 1
	prev1 := 1
	for i := 3; i <= N; i++ {
		prev2, prev1 = prev1, prev2+prev1
	}
	return prev1
}
