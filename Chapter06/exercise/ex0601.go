package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{FirstName: firstName, LastName: lastName, Age: age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age}
}

func main() {
	p := MakePerson("John", "Doe", 30)
	fmt.Println((p))
	pp := MakePersonPointer("Jane", "Doe", 25)
	fmt.Println((pp))
}
