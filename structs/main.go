package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type location struct {
	longiture float64
	latitude  float64
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	location
}

func main() {
	webb := person{firstName: "Bob", lastName: "Fischer"}
	fmt.Println(webb)

	eleph := person{
		firstName: "Eleph",
		lastName:  "Webb",
		contactInfo: contactInfo{
			email:   "elephwebb@trailblazers.com",
			zipCode: 94000,
		},
		location: location{
			latitude:  40.73,
			longiture: -75.93,
		},
	}

	/**
	* &variable: Give me the memory address of the value this variable is point at
	* *pointer: Give me the value this memory address is pointing at
	**/
	//elephPointer := &eleph
	//elephPointer.updateName("hpele")
	// same as above, in Go we can use
	// Go is a pass as value language
	eleph.updateName("PassingNameNotAsPointe")
	eleph.print()
}

/**
*   *person: this is a type description - it means we're working with a pointer to a person
*	*pointerToPerson: this is an operator - it means we want to manipulate the value the pointer is referencing
 */
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
