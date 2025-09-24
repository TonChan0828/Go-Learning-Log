package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{FirstName: "First", LastName: "Last", Age: i})
	}
}
