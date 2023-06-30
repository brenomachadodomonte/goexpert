package main

import "fmt"

func main() {
	events := []string{"event1", "event2", "event3", "event4"}

	// "event2", "event3", "event4" (skip first item)
	fmt.Println(events[1:])

	// "event1", "event2" (get till the second item)
	fmt.Println(events[:2])

	// "event2", "event3" "event4"(remove first item)
	events = append(events[:0], events[1:]...)
	fmt.Println(events)
}
