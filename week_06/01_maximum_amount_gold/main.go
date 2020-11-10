package main

import (
	"fmt"
)

func main() {
	var w, n int
	_, _ = fmt.Scanf("%d", &w)
	_, _ = fmt.Scanf("%d", &n)

	weights := make([]int, n+1)

	for i := 1; i <= n; i++ {
		var input int
		_, _ = fmt.Scanf("%d", &input)

		weights[i] = input
	}

	fmt.Printf("%d", maximumGold(w, n, weights))
}

func maximumGold(w, n int, weights []int) int {
	values := make([][]int, n+1)
	for i := range values {
		values[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			values[i][j] = values[i-1][j]
			if weights[i] <= j {
				val := values[i-1][j-weights[i]] + weights[i]
				if values[i][j] < val {
					values[i][j] = val
				}
			}
		}
	}

	return values[n][w]
}
