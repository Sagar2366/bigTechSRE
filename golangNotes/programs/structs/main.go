package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	fmt.Println("Struct Program ....")

	p1 := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex123@gmail.com",
			zipcode: 24536,
		},
	}

	var p2 person
	p2.firstName = "Sagar"
	p2.lastName = "Utekar"
	p2.contact = contactInfo{email: "sagar1234@gmail.com", zipcode: 24536}

	p1Pointer := &p1
	p1Pointer.updateFirstName("Allu")
	p1.printDetails()
	p2.printDetails()
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) printDetails() {
	fmt.Printf("Name : %+v\n", p)
}
