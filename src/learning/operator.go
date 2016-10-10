package main
import "fmt"

func main()  {
  var nilai = (((2+6)%3)* 4 - 2) / 3
  var isEqual = (nilai == 2)

  fmt.Println("jawabnya", isEqual)

  // boolean
  var left  = false
  var right = true

  var leftAndRight = left && right
  fmt.Printf("left and right \t (%t) \n", leftAndRight)

  var leftOrRight = left || right
  fmt.Printf("left or right \t (%t) \n", leftOrRight)

  var reverseLeft = !left
  fmt.Printf("!left \t (%t) \n", reverseLeft)

  // conditional
  var point = 3
  if point == 10 {
    fmt.Println("pass with excellence")
  } else if point > 5 {
    fmt.Println("pass only")
  } else if point == 4 {
    fmt.Println("almost pass")
  } else {
    fmt.Println("not pass. your value is \n", point)
  }

  // temporary variable
  var value = 8840.0
  if percent := value / 100; percent >=100 {
    fmt.Printf("%.1f%s good \n", percent, "%")
  } else if percent >= 70 {
    fmt.Printf("%.1f%s good \n", percent, "%")
  } else {
    fmt.Printf("%.1f%s not bad \n", percent, "%")
  }

  // switch
  var val = 5
  switch val {
  case 8:
    fmt.Println("perfect")
  case 7:
    fmt.Println("awesome")
  case 6:
    fmt.Println("enough")
  default:
    fmt.Println("not bad")
  }

  // multiple cases
  var score = 6
  switch score {
  case 8:
    fmt.Println("perfect")
  case 7,6,5,4:
    {
      fmt.Println("keep up the good work")
      fmt.Println("nice!")
    }
  default:
    fmt.Println("good")
  }
}
