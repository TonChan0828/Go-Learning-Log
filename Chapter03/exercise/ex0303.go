package main

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	emp1 := Employee{"John", "Doe", 1}
	emp2 := Employee{
		firstName: "Jane",
		lastName:  "Doe",
		id:        2,
	}
	var emp3 Employee
	emp3.firstName = "Max"
	emp3.lastName = "Smith"
	emp3.id = 3

	println(emp1.firstName, emp1.lastName, emp1.id)
	println(emp2.firstName, emp2.lastName, emp2.id)
	println(emp3.firstName, emp3.lastName, emp3.id)
}
