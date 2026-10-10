package main

import "fmt"

func main() {
	birds := []int{2, 5, 1, 0, 3, 1, 4, 2, 1, 0, 2, 3, 1, 1}
	fmt.Println("Total:", TotalBirdCount(birds))
	fmt.Println("Week 1:", BirdsInWeek(birds, 1))
	fmt.Println("Fixed:", FixBirdCountLog(birds))
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for _, bird := range birdsPerDay {
		sum += bird
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sum int
	for i := (week - 1) * 7; i < 2*week; i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay

}
