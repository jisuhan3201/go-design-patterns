package main

import (
	"fmt"

	"github.com/jisuhan3201/go-design-patterns/solid/srp"
)

func main() {
	j := srp.Journal{}
	j.AddEntry("first journal")
	fmt.Println(j.String())
}
