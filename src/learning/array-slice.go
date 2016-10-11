package main
import "fmt"

func main()  {
  var fruits = []string {"apple","grape","melon","pinneapple"}
  fmt.Println(fruits[0]) // "apple"

  var newFruits = fruits[0:2]
  fmt.Println(newFruits)
}
