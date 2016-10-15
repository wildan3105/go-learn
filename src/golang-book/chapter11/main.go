package main

import "fmt"
import "golang-book/chapter11/math"

// Finds the average of a series of 4 numbers defined
func main()  {
  xs := []float64{1,2,3,4}
  avg := math.Average(xs)
  fmt.Println(avg)
}
