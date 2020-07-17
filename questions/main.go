package main

type Naming interface {
	SetName(name string)
}

type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	p.name = name
}

type Student struct {
	Person
}

func main() {
	var (
		naming Naming
		student Student
	)
	naming = &student
	naming.SetName("some name")
}
