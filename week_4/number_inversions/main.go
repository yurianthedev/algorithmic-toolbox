package main

import (
	"fmt"
	"math"
)

var inversions int

func main() {
	var n int
	var seq []int

	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var val int
		_, _ = fmt.Scan(&val)
		seq = append(seq, val)
	}

	_ = numberOfInversions(seq)
	fmt.Println(inversions)
}

func numberOfInversions(seq []int) []int {
	if len(seq) <= 1 {
		return seq
	}

	half := int(math.Floor(float64(len(seq)) / 2))
	lower := numberOfInversions(seq[:half])
	upper := numberOfInversions(seq[half:])

	sorted, localInversionsCount := merge(lower, upper)
	inversions += localInversionsCount

	return sorted
}

func merge(lower, upper []int) ([]int, int) {
	var sorted []int
	var i, j, inversions int

	for i < len(lower) && j < len(upper) {
		if lower[i] <= upper[j] {
			sorted = append(sorted, lower[i])
			i++
		} else {
			sorted = append(sorted, upper[j])
			inversions += len(lower) - i
			j++
		}
	}

	sorted = append(sorted, lower[i:]...)
	sorted = append(sorted, upper[j:]...)

	return sorted, inversions
}
