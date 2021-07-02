package DP

func FibMemo(n int) int {
	m := make(map[int]int)
	return fibMemoHandler(n, m)
}

func fibMemoHandler(n int, memo map[int]int) int {

	if val, ok := memo[n]; ok {
		return val
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	memo[n] = fibMemoHandler(n-1, memo) + fibMemoHandler(n-2, memo)
	return memo[n]
}

func FibTab(n int) int {
	s := make([]int, n+1)
	s[0] = 0
	s[1] = 1

	for i := 2; i < len(s); i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]
}
