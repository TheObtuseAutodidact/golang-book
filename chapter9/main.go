// Problems
//
// What's the difference between a method and a function?

//    method has an explicit receiver allowing for dot
//    notation of the form thing.method()
//
// Why would you use an embedded anonymous field instead of a normal named field?

//    embedded field imply relation ship of the has-a or
//    is-a type
//
// Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.

package main

import "fmt"
import "math"

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

//func rectangleArea(x1, y1, x2, y2 float64) float64 {
//  l := distance(x1, y1, x1, y2)
//  w := distance(x1, y1, x2, y1)
//  return l * w
//}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	fmt.Println("2 sides of length", l)
	fmt.Println("2 sides of length", w)
	return 2.0*l + 2.0*w
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//func circleArea(c Circle) float64 {
//  return math.Pi * c.r*c.r
//}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

type Circle struct {
	x, y, r float64
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.perimeter()
	}
	return peri
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {

	person := Person{"Lenny"}
	person.Talk()

	andy := new(Android)
	andy.Name = "Andy"
	andy.Talk()

	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	rect := Rectangle{rx1, ry1, rx2, ry2}

	fmt.Println(rect.area())

	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	fmt.Printf("A circle with radius of %v has an area of %v\n", c.r, c.area())

	fmt.Println(totalArea(&c, &rect))
	fmt.Println(totalPerimeter(&rect))
	fmt.Println(totalPerimeter(&c))
	fmt.Println(totalPerimeter(&rect, &c))

}
