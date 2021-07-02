package DP

func CanSumMemo(t int, numbers []int) bool {
	m := make(map[int]bool)
	return canSumMemoHandler(t, numbers, m)
}

func canSumMemoHandler(targetSum int, candidates []int, memo map[int]bool) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	if v, ok := memo[targetSum]; ok {
		return v
	}

	for _, el := range candidates {

		remainder := targetSum - el
		if canSumMemoHandler(remainder, candidates, memo) {
			memo[targetSum] = true
			return true

		}
	}
	memo[targetSum] = false
	return memo[targetSum]

}

func CanSumTab(target int, candidates []int) bool {
	s := make([]bool, target+1)

	s[0] = true

	for i := 0; i < len(s); i++ {
		if s[i] == true {
			for _, num := range candidates {
				if i+num < len(s) {
					s[i+num] = true
				}
			}
		}
	}

	return s[target]
}
