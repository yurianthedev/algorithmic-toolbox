package main

import (
	"fmt"
)

func main() {
	var n int
	var values []int
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Scan(&v)
		values = append(values, v)
	}

	switch majority(values, 0, len(values)-1) {
	case 0:
		fmt.Println(0)
	default:
		fmt.Println(1)
	}
}

func majority(values []int, low, high int) int {
	if low == high {
		return values[low]
	}
	if low+1 == high {
		if values[low] != values[high] {
			return 0
		} else {
			return values[low]
		}
	}

	half := low + (high-low)/2
	x := majority(values, low, half)
	y := majority(values, half+1, high)

	var xCount, yCount int

	for _, value := range values[low : high+1] {
		if value == x {
			xCount++
		} else if value == y {
			yCount++
		}
	}

	n := (high + 1 - low) / 2
	if xCount > n {
		return x
	}
	if yCount > n {
		return y
	}

	return 0
}
