package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

type event struct {
	coordinate int
	eventType  rune
	index      int
}

func main() {
	var i int
	for {
		rand.Seed(time.Now().UnixNano())
		s := rand.Intn(200-100) + 100
		p := rand.Intn(300-100) + 100
		var starts, ends, points []int

		starts, ends = getRandStartsEnds(s)
		points = getRandSlice(p)

		payOffs := getPayoff(starts, ends, points)
		slowPayOff := getSlowPayOff(starts, ends, points)

		if !reflect.DeepEqual(payOffs, slowPayOff) {
			fmt.Printf("Case:\n%v\n%v\nPoints: %v\n", starts, ends, points)

			fmt.Printf("Rapid implementation: %v\n", payOffs)
			fmt.Printf("Slow implementation: %v\n", slowPayOff)

			panic("Different results")
		}

		fmt.Printf("Passed %d\n", i)
		i++
	}
}

func getRandSlice(s int) []int {
	rand.Seed(time.Now().UnixNano())
	segments := make([]int, 0)
	for i := 0; i < s; i++ {
		segments = append(segments, rand.Intn(500)-rand.Intn(500))
	}

	return segments
}

func getRandStartsEnds(n int) ([]int, []int) {
	rand.Seed(time.Now().UnixNano())
	var starts, ends []int

	for i := 0; i < n; i++ {
		starts = append(starts, rand.Intn(500)-rand.Intn(500))
		ends = append(ends, starts[i]+rand.Intn(250))
	}

	return starts, ends
}

func getSlowPayOff(starts, ends, points []int) []int {
	count := make([]int, len(points))

	for i, point := range points {
		for j := range starts {
			if point >= starts[j] && point <= ends[j] {
				count[i]++
			}
		}
	}

	return count
}

func getPayoff(starts, ends, points []int) []int {
	var events []event
	var segmentsCount int
	count := make([]int, len(points))

	for i := 0; i < len(starts); i++ {
		events = append(events, event{coordinate: starts[i], eventType: 's', index: i})
	}
	for i, value := range points {
		events = append(events, event{coordinate: value, eventType: 'p', index: i})
	}
	for i, value := range ends {
		events = append(events, event{coordinate: value, eventType: 'e', index: i})
	}
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].coordinate < events[j].coordinate
	})

	for _, event := range events {
		switch event.eventType {
		case 's':
			segmentsCount++
		case 'e':
			segmentsCount--
		case 'p':
			count[event.index] = segmentsCount
		default:
			panic("No recognized event type")
		}
	}

	return count
}
