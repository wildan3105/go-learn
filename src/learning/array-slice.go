package main
import "fmt"

func main()  {
  var fruits = []string {"apple","grape","melon","pinneapple"}
  fmt.Println(fruits[0]) // "apple"

  var newFruits = fruits[:]
  fmt.Println(newFruits)

  // old and new slice
  var buah = []string{"apple","grape","pine","orange"}
  var aBuah = buah[0:3]
  var bBuah = buah[1:4]

  var aaBuah = aBuah[1:2]
  var bbBuah = bBuah[0:1]

  fmt.Println(buah)
  fmt.Println(aBuah)
  fmt.Println(bBuah)
  fmt.Println(aaBuah)
  fmt.Println(bbBuah)

  //baBuah[0] = "durian"

  fmt.Println("length of buah : ", len(buah))
  fmt.Println()

  // append
  var cinemas = []string {"marvel","green","arrow"}
  var cCinemas = append(cinemas, "hobbit")
  fmt.Println(cinemas)
  fmt.Println(cCinemas)

  // copy
  var colors = []string{"green"}
  var aColors = []string{"red","blue"}

  var copiedElement = copy(colors, aColors)
  fmt.Println(colors)
  fmt.Println(aColors)
  fmt.Println(copiedElement)
  
}
