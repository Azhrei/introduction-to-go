package main

import "fmt"

type Point struct{ X, Y int }

func (p Point) String() string { return fmt.Sprintf("point: x=%d, y=%d", p.X, p.Y) }

// START OMIT
func (p *Point) Move(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{X: 20, Y: 20}
	fmt.Println(p)

	p.Move(20, 10)	// This one now works!
	fmt.Println(p)
}
// END OMIT
