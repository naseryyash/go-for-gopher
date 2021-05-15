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

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// below is a shortcut for the above 2 commented lines
	jim.updateName("Jimmy")
	jim.print()

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "Bill"
	fmt.Println(&name)

	updateString(name)
	fmt.Println(name)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func updateString(s string) {
	s = "Shrek"
}
