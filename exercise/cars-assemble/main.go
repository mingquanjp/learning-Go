package main

import "fmt"

func main() {
	fmt.Println("Per hour:", CalculateWorkingCarsPerHour(100, 80))
	fmt.Println("Per minute:", CalculateWorkingCarsPerMinute(100, 80))
	fmt.Println("Cost:", CalculateCost(37))
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	first, second := carsCount/10, carsCount%10
	return uint(first)*95000 + uint(second)*10000
}
