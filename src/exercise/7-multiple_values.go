package main
import "fmt"

func main()  {
  x, y := f()
  fmt.Println(x,y)
  fmt.Println(add(1,2,3))
}

func f() (int, int){
  return 5,6
}

func add(args ...int) int {
  total := 0
  for _, v := range args{
    total += v
  }
  return total
}
