package main
import "fmt"

func main()  {
  fmt.Println("Enter degree in fahrenheit : ")
  var input float64
  fmt.Scanf("%f", &input)
  var output float64 = ((input - 32) * 5/9)
  fmt.Println("result in celcius : %.2f", output)
}
