package main

import "fmt"

// Circle is a datatype that representes a circle
type Circle struct {
	radius float32
}

// Square is a datatype that representas a square
type Square struct {
	side float32
}

// Area is a function that evaluates the area, attatch to a circle
func (c Circle) Area() float32 {
	return c.radius * c.radius * 3.1416
}

// Area is a function that is attached to a square
func (s Square) Area() float32 {
	return s.side * s.side
}

// Area is an interface that just neeads an area function
type Area interface {
	Area() float32
}

// Meassure is a function o ann interface Area
func Meassure(a Area) float32 {
	return a.Area()
}

func main() {
	circle := Circle{2}
	square := Square{4}

	fmt.Println("with methods")
	fmt.Println("area of a circle with radius 2 is: ", circle.Area())
	fmt.Println("area of a square with side 4 is: ", square.Area())

	fmt.Println("with interfaces")
	fmt.Println(Meassure(circle))
	fmt.Println(Meassure(square))
}
