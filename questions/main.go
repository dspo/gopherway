package main

import (
	"fmt"
)

type Naming interface {
	SetName(name string)
	GetName() string
}

type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) GetName() string {
	return p.name
}

type Student struct {
	*Person
}

func main() {
	var (
		naming Naming
		student Student
	)
	naming = student
	naming.SetName("some name")
	fmt.Println(naming.GetName())
}
