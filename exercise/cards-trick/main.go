package main

import "fmt"

func main() {
	items := []int{1, 2, 3}
	fmt.Println("Favorite cards:", FavoriteCards())
	fmt.Println("Item 1:", GetItem(items, 1))
	fmt.Println("Prepend:", PrependItems(items, 8, 9))
}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	FavoriteCards := []int{2, 6, 9}
	return FavoriteCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}

	return slice[index]
}

// SetItem writes an item to a slice at a given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	slice = append(values, slice...)
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}
