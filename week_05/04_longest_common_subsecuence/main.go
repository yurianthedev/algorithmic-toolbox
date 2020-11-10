package main

import "fmt"

func main() {
	var m int
	_, _ = fmt.Scan(&m)
	s1 := make([]int, m)
	for i := 0; i < m; i++ {
		var value int
		_, _ = fmt.Scan(&value)
		s1[i] = value
	}

	var n int
	_, _ = fmt.Scan(&n)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		var value int
		_, _ = fmt.Scan(&value)
		s2[i] = value
	}

	fmt.Printf("%d", longestCommonSubsequence(m, n, s1, s2))

}

func longestCommonSubsequence(m, n int, s1, s2 []int) int {
	computations := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		computations[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i == 0 {
				continue
			}
			if s1[i-1] == s2[j-1] {
				computations[i][j] = computations[i-1][j-1] + 1
			} else if computations[i-1][j] >= computations[i][j-1] {
				computations[i][j] = computations[i-1][j]
			} else {
				computations[i][j] = computations[i][j-1]
			}
		}
	}

	return computations[m][n]
}
