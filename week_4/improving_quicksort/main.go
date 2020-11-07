package main

import (
	"fmt"
	"io"
	"math/rand"
)

func main() {
	var n int
	var values []int
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var value int
		_, _ = fmt.Scan(&value)
		values = append(values, value)
	}

	_, _ = io.Reader(nil).Read(nil)

	quickSort(values, 0, len(values)-1)
	for _, val := range values {
		fmt.Printf("%d ", val)
	}
}

func quickSort(values []int, low, high int) {
	if low >= high {
		return
	}

	random := rand.Intn(high-low) + low
	values[low], values[random] = values[random], values[low]
	m1, m2 := partition3(values, low, high)
	quickSort(values, low, m1-1)
	quickSort(values, m2+1, high)
}

func partition3(values []int, low, high int) (int, int) {
	x := values[low]
	lt := low
	gt := high
	i := low + 1

	for i <= gt {
		if x > values[i] {
			values[i], values[lt] = values[lt], values[i]
			lt++
			i++
		} else if x < values[i] {
			values[i], values[gt] = values[gt], values[i]
			gt--
		} else {
			i++
		}
	}

	return lt, gt
}
