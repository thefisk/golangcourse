package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo  // can also just use the struct name as the person property, so could just type "contactInfo"
}

func main() {
	jim := person {
		firstName: "Jim",
		lastName: "Smith",
		contact: contactInfo {
			email: "jim@gmail.com",
			zipCode: 12345,
		},
	}

	// & syntax to create a pointer
	pointerToJim := &jim
	// First function will update jim's memory address
	pointerToJim.updateNameByPointer("jonny")
	// Second method won't update jim
	jim.updateNameByValue("jimmy")
	// However, if a function receives a pointer, you don't have to explicitly create a pointer
	jim.updateNameByPointer("jansen")
	jim.print()
	// %p on Printf shows the Pointer value - showing that it stores a memory address
	fmt.Printf("\npointerToJim is memory address %p",pointerToJim)
}


// This function doesn't update the 'jim' instance because Go passes by value, so 'p' is a copy of 'jim'
func (p person) updateNameByValue(newFirstName string) {
	p.firstName = newFirstName
}

// Instead, to update 'jim' we need to pass in a pointer, using * syntax
func (p *person) updateNameByPointer(newFirstName string) {
	// The * reads the contents of the memory at the pointer
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%v", p)
}