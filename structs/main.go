package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Henson",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipcode: 87114,
		},
	}

	fmt.Printf("%+v", jim)
}