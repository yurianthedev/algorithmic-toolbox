package main

import (
	"fmt"
	"math"
)

func main() {
	var string1, string2 string
	_, _ = fmt.Scan(&string1)
	_, _ = fmt.Scan(&string2)

	fmt.Printf("%d", editDistance(string1, string2))

}

func editDistance(string1, string2 string) int {
	n, m := len(string1), len(string2)

	computations := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		computations[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			if i == 0 {
				computations[i][j] = j
			} else if j == 0 {
				computations[i][j] = i
			} else if string1[i-1] == string2[j-1] {
				computations[i][j] = computations[i-1][j-1]
			} else {
				computations[i][j] = 1 + int(math.Min(math.Min(float64(computations[i][j-1]), float64(computations[i-1][j])), float64(computations[i-1][j-1])))
			}
		}
	}

	return computations[n][m]
}
