package main

import "fmt"

type Dragon struct {
	Bird
	Lizard
}

type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying!!")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age >= 10 {
		fmt.Println("Crawling!!")
	}
}

func main() {
	d := Dragon{}
	// d.Age = 10 // Ambiguous selector error
	d.Bird.Age = 10
	d.Lizard.Age = 10 // 하나의 객체에 두번 할당하는 문제
	d.Fly()
	d.Crawl()
}
