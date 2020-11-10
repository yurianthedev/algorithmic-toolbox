package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	if n < 3 {
		fmt.Printf("%d", 0)
	}

	souvenirs := make([]int, n+1)
	var sum int

	for i := 1; i <= n; i++ {
		var input int
		_, _ = fmt.Scanf("%d", &input)

		sum += input
		souvenirs[i] = input
	}

	if sum%3 == 0 {
		fmt.Printf("%d", couldPartition(sum/3, souvenirs))
	} else {
		fmt.Printf("%d", 0)
	}
}

func couldPartition(W int, souvenirs []int) int {
	values := make([][]int, len(souvenirs))
	for i := range values {
		values[i] = make([]int, W+1)
	}

	var count int

	for i := 1; i <= len(souvenirs)-1; i++ {
		for j := 1; j <= W; j++ {
			values[i][j] = values[i-1][j]
			if souvenirs[i] <= j {
				val := values[i-1][j-souvenirs[i]] + souvenirs[i]
				if values[i][j] < val {
					values[i][j] = val
				}
			}
			if values[i][j] == W {
				count++
			}
		}
	}

	if count < 3 {
		return 0
	}

	return 1
}
