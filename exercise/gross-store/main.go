package main

import "fmt"

func main() {
	units := Units()
	bill := NewBill()
	AddItem(bill, units, "eggs", "dozen")
	quantity, ok := GetItem(bill, "eggs")
	fmt.Println("Bill:", bill, "eggs:", quantity, "found:", ok)
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := make(map[string]int)
	measurements["quarter_of_a_dozen"] = 3
	measurements["half_of_a_dozen"] = 6
	measurements["dozen"] = 12
	measurements["small_gross"] = 120
	measurements["gross"] = 144
	measurements["great_gross"] = 1728
	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := units[unit]; !exists {
		return false
	}
	bill[item] = units[unit]
	if res := bill[item]; res != 0 {
		res += units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := bill[item]; !exists {
		return false
	}
	if _, exists := units[unit]; !exists {
		return false
	}
	newQuantity := bill[item] - units[unit]
	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
	}
	bill[item] = newQuantity

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, exists := bill[item]; !exists {
		return 0, false
	}
	return bill[item], true
}
