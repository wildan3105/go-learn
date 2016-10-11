package main
import "fmt"

func main()  {
  // Number, int
  fmt.Println("1+1 = ", 1+1)
  fmt.Println("1+1 = ", 1.0+1.0)

  // String
  fmt.Println(len("Hello world"))
  fmt.Println("Hello world"[1])
  fmt.Println("Hello"+"world")

  // Booleans
  fmt.Println(true && true) // true
  fmt.Println(true && false) // false
  fmt.Println(true || true) // true
  fmt.Println(!true) // false
}
