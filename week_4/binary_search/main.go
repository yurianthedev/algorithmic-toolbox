package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	_, _ = fmt.Scan(&n)
	var a []int

	for i := 0; i < n; i++ {
		var element int
		_, _ = fmt.Scan(&element)
		a = append(a, element)
	}

	_, _ = fmt.Scan(&k)
	for i := 0; i < k; i++ {
		var element int
		_, _ = fmt.Scan(&element)

		fmt.Printf("%d ", searchValueIndex(element, 0, len(a)-1, a))
	}
}

func searchValueIndex(value int, low int, high int, a []int) int {
	if high < low {
		return -1
	}

	half := int(math.Floor(float64(low + ((high - low) / 2))))
	switch {
	case value == a[half]:
		return half
	case value < a[half]:
		return searchValueIndex(value, low, half-1, a)
	default:
		return searchValueIndex(value, half+1, high, a)
	}
}
