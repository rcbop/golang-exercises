package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

// main
func main() {
	var alex person
	var contact contactInfo
	contact.email = "alex.anderson@gmail.com"
	contact.zipCode = 5555555
	// alex := person{
	// 	 firstName: "Alex",
	//   lastName: "Anderson",
	//   contact: contactInfo {
	//	 	 email: "jim@gmail.com",
	//  	 zipCode:90000,
	//   }
	// },
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo = contact
	// alexPointer := &alex
	// alexPointer.updateName("alexy")
	alex.updateName("alexy")
	alex.print()
}
