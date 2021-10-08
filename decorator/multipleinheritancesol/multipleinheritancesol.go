package main

import "fmt"

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}
func (b *Bird) SetAge(age int) {
	b.age = age
}
func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}
func (l *Lizard) SetAge(age int) {
	l.age = age
}
func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!!")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

// 여전히 inconsistency가 존재 한다. (d.lizard.age 가 다를 경우)
// 하지만 d.lizard.age를 private field로 하여금 접근하지 못하게 한다.
func (d *Dragon) Age() int {
	return d.bird.age
}
func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}
func (d *Dragon) Fly() {
	d.bird.Fly()
}
func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}
func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

func main() {
	d := NewDragon()
	d.SetAge(10)
	d.Fly()
	d.Crawl()
}
