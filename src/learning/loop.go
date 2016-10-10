package main
import "fmt"

func main()  {
  // ONLY FOR
  for i:=0; i < 5; i++ {
    fmt.Println("Angka ",i)
  }

  // with condition
  var it = 0
  for it < 5 {
    fmt.Println("Number ",it)
    it++
  }

  // without argument/condition
  var j = 0
  for {
    fmt.Println("Angka ", j)
    j++

    if j == 7 {
      break
    }
  }

  // the use of CONTINUE and break
  for k:=1; k <= 10; k++ {
    if k % 2 == 1 {
      continue
    }
    if k > 8 {
      break
    }

    fmt.Println("Number ", k)
  }

  // nested loop
  for m:=0; m < 5; m++ {
    for n:=m; n < 5; n++ {
      fmt.Print(n, " ")
    }
    fmt.Println()
  }

  // label
  outerLoop:
  for o:= 0; o < 5; o++{
    for p:= 0; p < 5; p++{
      if o == 3 {
        break outerLoop
      }
      fmt.Print("matriks [", o, "][", p, "]", "\n" )
    }
  }
}
