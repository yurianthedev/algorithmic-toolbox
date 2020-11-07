package main

import (
	"fmt"
	"math"
	"sort"
)

type point struct {
	x, y int
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	points := make([]point, n)

	for i := 0; i < n; i++ {
		var x, y int
		_, _ = fmt.Scan(&x)
		_, _ = fmt.Scan(&y)

		points[i] = point{x, y}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})
	Y := make([]point, n)
	copy(Y, points)
	sort.Slice(Y, func(i, j int) bool {
		return Y[i].y < Y[j].y
	})
	fmt.Printf("%v\n%v\n", points, Y)

	d := closestDistance(points, Y)
	fmt.Printf("%.4f", d)
}

func closestDistance(X, Y []point) float64 {
	if len(X) == 2 {
		return calcDistance(X[0], X[1])
	} else if len(X) == 3 {
		return math.Min(math.Min(calcDistance(X[0], X[1]), calcDistance(X[1], X[2])), calcDistance(X[0], X[2]))
	}

	half := int(math.Floor(float64(len(X)) / 2))
	pointsLeft := X[:half]
	pointsRight := X[half:]

	XL, XR := pointsLeft, pointsRight
	YL, YR := sortedSplit(pointsLeft, pointsRight, Y)

	minLeft := closestDistance(XL, YL)
	minRight := closestDistance(XR, YR)

	sigma := math.Min(minLeft, minRight)
	sigmaPrime := closestSplitPoint(Y, pointsRight[0], sigma)

	return math.Min(sigma, sigmaPrime)
}

func closestSplitPoint(Y []point, mid point, sigma float64) float64 {
	YPrime := make([]point, 0)

	for i, point := range Y {
		if i == 0 {
			continue
		}
		if calcDistance(point, mid) > 2*sigma {
			YPrime = append(YPrime, point)
		}
		if i == 6 {
			break
		}
	}

	for i, point := range Y {
		for j, yPoint := range YPrime {
			if i == 0 || i != j+1 {
				tempSigma := calcDistance(point, yPoint)
				if tempSigma < sigma {
					sigma = tempSigma
				}
			}
		}
	}

	return sigma
}

func sortedSplit(pointsLeft, pointsRight, Y []point) ([]point, []point) {
	left := make([]point, 0, len(pointsLeft))
	right := make([]point, 0, len(pointsRight))

	for _, point := range Y {
		if point.x < pointsRight[0].x {
			left = append(left, point)
		} else {
			right = append(right, point)
		}
	}

	return left, right
}

func calcDistance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
}
