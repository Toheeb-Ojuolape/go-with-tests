package structs

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Breadth * r.Length
}

func (c Circle) Area() float64 {
	return (22 / 7) * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Breadth + rectangle.Length)
}
