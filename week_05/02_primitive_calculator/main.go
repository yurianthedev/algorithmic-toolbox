package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	minOperations, steps := calculate(n)

	fmt.Println(minOperations)
	for i := len(steps) - 2; i >= 0; i-- {
		fmt.Printf("%d ", steps[i])
	}
}

func calculate(value int) (int, []int) {
	computations := make([]int, value+1)

	for i := range computations {
		if i < 2 {
			continue
		}

		multiplyBy2, multiplyBy3 := math.MaxInt64, math.MaxInt64
		sum := computations[i-1] + 1
		if i%2 == 0 {
			multiplyBy2 = computations[i/2] + 1
		}
		if i%3 == 0 {
			multiplyBy3 = computations[i/3] + 1
		}

		computations[i] = int(math.Min(math.Min(float64(sum), float64(multiplyBy2)), float64(multiplyBy3)))
	}

	minComp := computations[value]
	steps := make([]int, 0, minComp+1)
	steps = append(steps, value)

	for i := value; i >= 1; {
		if i%3 == 0 && computations[i]-1 == computations[i/3] {
			steps = append(steps, i/3)
			i /= 3
		} else if i%2 == 0 && computations[i]-1 == computations[i/2] {
			steps = append(steps, i/2)
			i /= 2
		} else {
			steps = append(steps, i-1)
			i--
		}
	}

	return minComp, steps
}
