package main
import (
  "fmt"
  "time"
  "math/rand"
)


// give some delay

func f(n int){
  for i:=0; i<10; i++{
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(1250))
    time.Sleep(time.Millisecond * amt)
  }
}

func main()  {
  for i:=0; i<10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
