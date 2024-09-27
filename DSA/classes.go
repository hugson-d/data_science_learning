package main

import (
	"fmt"
)

type Robot struct {
	Name  string
	Age   int
	color string
}

// method with a pointer receiver (similar to a class method)

func (r *Robot) ChangeName(newName string) {
	r.Name = newName
}

func (r *Robot) Greet() {
	fmt.Printf("Hello, my name is %s.\n", r.Name)
}

func (r *Robot) setAge(Age int) {
	r.Age = Age
}

func classes_test() {
	person := &Robot{Name: "John Doe", Age: 30, color: "blue"}

	person.Greet()

	person.ChangeName("Jane Doe")
	person.Greet()

	person.setAge(25)
	fmt.Printf("I am %d years old.\n", person.Age)

	person.Greet()
}
