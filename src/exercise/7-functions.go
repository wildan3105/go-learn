package main
import "fmt"

func average(xs []float64) float64 {
  total := 0.0

  for _, v:= range xs {
    total += v
  }
  return total / float64(len(xs))
}

func main()  {
  xs := []float64{10, 20, 30, 40}
  fmt.Println(average(xs))

  // another example
  fmt.Println(f1)
}

func f1() int {
  return f2()
}

func f2() int {
  return 20
}
