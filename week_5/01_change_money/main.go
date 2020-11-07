package main

import (
	"fmt"
	"math"
)

func main() {
	var money int
	_, _ = fmt.Scan(&money)
	denominations := [...]int{1, 3, 4}

	fmt.Printf("%d", minNumberCoins(money, &denominations))
}

func minNumberCoins(money int, denominations *[3]int) int {
	computations := make([]int, money+1)
	for i := range computations {
		if i == 0 {
			continue
		}

		computations[i] = math.MaxInt64

		for _, denomination := range denominations {
			if denomination <= i {
				numberOfCoins := computations[i-denomination]
				if numberOfCoins+1 < computations[i] {
					computations[i] = numberOfCoins + 1
				}
			}
		}
	}

	return computations[len(computations)-1]
}
