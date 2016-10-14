package main
import (
  "fmt";
  "math"
)


// hardcoding

/*
func rectangleArea(x1, y1, x2, y2 float64) float64 {sr
  l := distance(x1,y1,x1,y2)
  w := distance(x1,y1,x2,y1)
  return l*w
}

func circleArea(x,y,r float64) float64{
  return math.Pi * r*r
}

func main(){
  var rx1, ry1 float64 = 0, 0
  var rx2, ry2 float64 = 10, 10
  var cx, cy, cr float64 = 0, 0, 5
  fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
  fmt.Println(circleArea(cx, cy, cr))
}
*/

// better approach with struct
type Circle struct {
  x, y, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

// main function
func main(){
  c := Circle{0, 0, 5}
  r := Rectangle{0,0,10,10}
  fmt.Println("Area of circle is", c.area())
  fmt.Println("Area of rectangle is ", r.area())
  a := new(Android)
  a.Person.Talk()
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l*w
}

// Embedded types
type Person struct {
  Name string
}

func(p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
  Person Person
  Model string
}
