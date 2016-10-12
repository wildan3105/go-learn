package main
import "fmt"

var x int = 1

func main(){
  if x % 2 == 0 {
    y := true
    fmt.Println("1,",y)
  } else {
    y := false
    fmt.Println("0,",y)
  }
}
