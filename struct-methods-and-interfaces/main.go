package structmethodsandinterfaces

import "math"

// Shape is a abstration to any shape
type Shape interface {
	Area() float64
}

// Rectangle is a type of shape
type Rectangle struct {
	Widht  float64
	Height float64
}

// Area is a Method to calc area of retangule
func (r Rectangle) Area() float64 {
	return r.Widht * r.Height
}

// Circle is a type of shape
type Circle struct {
	Radio float64
}

// Area is a Method to calc area of Circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}

// Triangle is a type of shape
type Triangle struct {
	Base   float64
	Height float64
}

//Area is a Method to calc area of Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Perimeter calcs a path that encompasses/surrounds/outlines a shape or its length.
func Perimeter(retangulo Rectangle) float64 {
	return 2 * (retangulo.Widht + retangulo.Height)
}

func main() {

}
