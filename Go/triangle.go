package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}
type shape interface {
	area() float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}
func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
func printArea(h shape) {
	fmt.Println(h.area())
}
func main() {
	tria := triangle{height: 5, base: 6}
	sq := square{sideLength: 7}
	printArea(tria)
	printArea(sq)
}
