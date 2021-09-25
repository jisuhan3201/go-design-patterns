package main

import "fmt"

type Person struct {
	Name string
	Age  int
	// EyeCount int
}

// factory function
// func NewPerson(name string, age int) Person {
// 	return Person{name, age}
// }
func NewPerson(name string, age int) *Person {
	// EyeCount 값 처럼 Default 값을 주고 싶은 경우
	// return &Person{name, age, 2}
	// 객체 field를 Validation 하고 싶을 경우
	// if age < 16 {
	// 	//....
	// }
	return &Person{name, age}
}

func main() {
	// init directly
	p := Person{"John", 22}
	fmt.Println(p)

	// use a constructor
	p2 := NewPerson("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)
}
