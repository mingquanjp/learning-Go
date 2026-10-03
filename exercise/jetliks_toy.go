package exercise

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery int
	batteryDrain int	
	speed int
	distance int
}


// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:100, 
		speed: speed,
		batteryDrain: batteryDrain,
		distance: 0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}


// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
// func Drive(car Car) Car {

// 	if car.battery == 0 {
// 		return car
// 	}
// 	car.battery-=car.batteryDrain
// 	car.distance += car.speed
// 	return car
// }

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	scale := float64(car.battery / car.batteryDrain)
	if float64(car.speed) * scale >= float64(track.distance){
		return true
	}
	return false

}




// TODO: define the 'Drive()' method
func (car Car) Drive() {
	if car.battery < car.batteryDrain {
		fmt.Println("Cannot drive")
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
}

func (car Car) DisplayDistance() string{
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string{
	return fmt.Sprintf("Battery at %d%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
	scale := car.battery / car.batteryDrain
	car.distance = car.speed * scale
	if car.distance >= trackDistance {
		return true
	}
	return false
}

// TODO: define the 'DisplayDistance() string' method

// TODO: define the 'DisplayBattery() string' method

// TODO: define the 'CanFinish(trackDistance int) bool' method

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.