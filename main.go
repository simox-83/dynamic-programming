package main

import (
	"dynamic-programming/DP"
	"fmt"
)

func main() {
	fmt.Println(DP.FibMemo(8))
	fmt.Println(DP.FibMemo(50))
	fmt.Println(DP.FibMemo(100))
	fmt.Println()
	fmt.Println(DP.FibTab(8))
	fmt.Println(DP.FibTab(50))
	fmt.Println(DP.FibTab(100))
	fmt.Println()
	fmt.Println(DP.GridTravelerMemo(1, 1))
	fmt.Println(DP.GridTravelerMemo(2, 3))
	fmt.Println(DP.GridTravelerMemo(3, 2))
	fmt.Println(DP.GridTravelerMemo(3, 3))
	fmt.Println(DP.GridTravelerMemo(18, 18))
	fmt.Println()
	fmt.Println(DP.GridTravelerTab(1, 1))
	fmt.Println(DP.GridTravelerTab(2, 3))
	fmt.Println(DP.GridTravelerTab(3, 2))
	fmt.Println(DP.GridTravelerTab(3, 3))
	fmt.Println(DP.GridTravelerTab(18, 18))
	fmt.Println()
	fmt.Println(DP.CanSumMemo(7, []int{2, 3}))
	fmt.Println(DP.CanSumMemo(7, []int{5, 3, 4, 7}))
	fmt.Println(DP.CanSumMemo(7, []int{2, 4}))
	fmt.Println(DP.CanSumMemo(8, []int{2, 3, 5}))
	fmt.Println(DP.CanSumMemo(300, []int{7, 14}))
	fmt.Println()
	fmt.Println(DP.CanSumTab(7, []int{2, 3}))
	fmt.Println(DP.CanSumTab(7, []int{5, 3, 4, 7}))
	fmt.Println(DP.CanSumTab(7, []int{2, 4}))
	fmt.Println(DP.CanSumTab(8, []int{2, 3, 5}))
	fmt.Println(DP.CanSumTab(300, []int{7, 14}))

}
