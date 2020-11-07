package main

import (
	"fmt"
	"sort"
)

type event struct {
	coordinate int
	eventType  rune
	index      int
}

func main() {
	var s, p int
	var starts, ends, points []int

	_, _ = fmt.Scan(&s)
	_, _ = fmt.Scan(&p)

	for i := 0; i < s; i++ {
		var a, b int
		_, _ = fmt.Scan(&a)
		_, _ = fmt.Scan(&b)

		starts = append(starts, a)
		ends = append(ends, b)
	}

	for i := 0; i < p; i++ {
		var value int
		_, _ = fmt.Scan(&value)

		points = append(points, value)
	}

	payOffs := getPayoff(starts, ends, points)
	for i, value := range payOffs {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(value)
	}
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
