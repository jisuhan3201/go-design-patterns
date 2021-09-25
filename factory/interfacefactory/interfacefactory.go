package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, I am %s and %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Hi, I am tired")
}

// return interface
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	// p := NewPerson("Jane", 20)
	p := NewPerson("Sam", 120)
	p.SayHello()
}
