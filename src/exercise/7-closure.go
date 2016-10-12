package main
import "fmt"

func main()  {
  // add:= func(x,y int ) int {
  //   return x+y
  // }
  // fmt.Println(add(2,3))

  // generate even numbers
  nextEven := makeGen()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())
}

func makeGen() func() uint{
  i := uint(0)
  return func() (ret uint){
    ret = i
    i += 2
    return
  }
}
