package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	quan := person{
		firstName : "Quan",
		lastName : "Nguyen",
		contact : contactInfo{
			email : "quan@gmail.com",
			zipCode : 12345,
		},
	}
	quan.print()
	quanPinter := &quan
	quanPinter.updateName("Quannn Nguyennn")
	quan.print()
	quan.updateName("John")
	quan.print()

	
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

