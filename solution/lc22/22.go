package lc22

func run(current string, n, left, right int) []string {
	if n == left && n == right {
		return []string{current}
	}
	if left == right {
		next := current + "("
		return run(next, n, left+1, right)
	}
	if left < n {
		leftResult, rightResult := run(current+"(", n, left+1, right), run(current+")", n, left, right+1)
		return append(leftResult, rightResult...)
	}
	return run(current+")", n, left, right+1)
}

func generateParenthesis(n int) []string {
	return run("", n, 0, 0)
}
