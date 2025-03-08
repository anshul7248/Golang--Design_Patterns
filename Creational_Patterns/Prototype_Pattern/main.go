package main

import "fmt"

type Prototype interface {
	Clone() Prototype
	GetDetails() string
}

type Employee struct {
	Name        string
	Age         int
	Designation string
}

func (e *Employee) Clone() Prototype {
	return &Employee{
		Name:        e.Name,
		Age:         e.Age,
		Designation: e.Designation,
	}
}

func (e *Employee) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d, Designtation: %s", e.Name, e.Age, e.Designation)
}

func main() {
	emp1 := &Employee{Name: "Anshul", Age: 25, Designation: "Golang Developer"}

	emp2 := emp1.Clone().(*Employee)

	emp2.Name = "Anil"

	fmt.Println("Original Employee Details--->>", emp1.GetDetails())
	fmt.Println("Clone Employee Details--->>", emp2.GetDetails())

}
