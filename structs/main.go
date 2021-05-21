package main

// Structs are like objects in JS, etc
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
}
