// structs2
// Make me compile!

package main

import "fmt"

type Phone struct {
	number string
}

type Person struct {
	// don't just create the phone field here. embed a new struct
	name  string
	age   int
	phone Phone
}

func main() {
	// contactDetails := ContactDetails{}
	person := Person{name: "John", age: 32, phone: Phone{number: "+1 234 56 78"}}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
