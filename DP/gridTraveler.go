package DP

import "fmt"

func GridTravelerMemo(m, n int) int {
	memo := make(map[string]int)
	return gridTravelerMemoHandler(m, n, memo)
}

func gridTravelerMemoHandler(m, n int, memo map[string]int) int {

	key := fmt.Sprintf("%d,%d", m, n)
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	if v, ok := memo[key]; ok {
		return v
	}

	memo[key] = gridTravelerMemoHandler(m, n-1, memo) + gridTravelerMemoHandler(m-1, n, memo)
	return memo[key]
}

func GridTravelerTab(m, n int) int {
	table := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		table[i] = make([]int, n+1)
	}

	table[1][1] = 1

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if j+1 <= n {
				table[i][j+1] += table[i][j]
			}
			if i+1 <= m {
				table[i+1][j] += table[i][j]
			}
		}
	}

	return table[m][n]
}
