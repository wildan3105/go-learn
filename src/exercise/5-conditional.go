package main
import "fmt"

func main()  {
  i := 1
  for i <=100 {
    if i % 3 ==0 {
      // genap
      fmt.Println(i)
    } else {
      // ganjil
      // fmt.Println("ganjil")
    }
    i++
  }
}
