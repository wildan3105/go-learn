package main
import "fmt"

func main()  {
  fmt.Println(fibonacci(6))
}

func fibonacci(x uint) uint {
  if x < 2 {
    return x
  }
  return fibonacci(x-1) + fibonacci(x-2)
}
