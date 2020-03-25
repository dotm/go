package shapes

import "math"

//Shape interface
type Shape interface {
	Area() float64
}

//Rectangle has width and height
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle has radius
type Circle struct {
	Radius float64
}

//Triangle has base and height
type Triangle struct {
	Base   float64
	Height float64
}

//Area get the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Area get the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Area get the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//Perimeter counts the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
