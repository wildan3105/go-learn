package main
import "fmt"

// func one(xPtr *int) {
//   *xPtr = 1
// }
//
// func main()  {
//   // x := 5
//   //zero(&x)
//   xPtr := new(int)
//   one(xPtr)
//   fmt.Println(*xPtr)
// }

func square(x *float64) {
*x = *x * *x
}
func main() {
x := 1.5
square(&x)

fmt.Println(x)
}
