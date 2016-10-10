package main
import "fmt"

func main()  {
  var point = 6
  switch {
  case point == 6:
    fmt.Println("perfect")
  case (point < 8) && (point > 3):
    fmt.Println("awesome")
  default:
    {
      fmt.Println("not bad")
      fmt.Println("you need to learn more")
    }
  }
  /*
    fallthrough : check to the next condition
  */
  var score = 6
  switch {
  case score == 8:
    fmt.Println("perfect")
  case (score < 8) && (score > 3):
    fmt.Println("awesome")
    fallthrough
  case score < 5:
    fmt.Println("you need to learn more")
  default:
    {
      fmt.Println("not bad")
      fmt.Println("nice enough")
    }
  }

  // nested conditional
//   var poin = 10
//   if poin > 7 {
//     switch poin {
//     case 10:
//       fmt.Println("perfect")
//     default:
//       fmt.Println("nice!")
//     }
//   }
//   else {
//     if poin == 5 {
//       fmt.Println("close enough")
//     } else if poin == 3 {
//       fmt.Println("keep trying")
//     }
//   else {
//     fmt.Println("you can do it!")
//     if poin == 0 {
//       fmt.Println("try harder")
//     }
//   }
// }

}
