package main

import (
	"fmt"
	"github.com/brenomachadodomonte/goexpert/6-packaging/1/math"
)

func main() {
	m := math.Math{A: 10, B: 20}
	fmt.Println("Hello World!")
	fmt.Printf("Sum: %d", m.Add())
}
