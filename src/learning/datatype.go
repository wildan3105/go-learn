package main
import "fmt"

func main()  {
  name:= new(string)

  fmt.Println(name)
  fmt.Println(*name)

  /*
  variable declaration using make
  only valid for these 3 vars : channel, slice, map
  */
  // numeric non-decimal variable
  var positiveNumb uint8 = 99
  var negativeNumb = -1212

  fmt.Printf("positive numb : %d \n", positiveNumb)
  fmt.Printf("negative numb : %d \n", negativeNumb)

  // bilangan cacah
  var decimalNumb = 2.62
  fmt.Printf("decimal numb : %f \n", decimalNumb)
  fmt.Printf("decimal numb : %.3f \n", decimalNumb)

  // boolean var
  var exist bool = true
  fmt.Printf("exist ? %t \n", exist)

  // string
  var msg = `My name is Wildan.
  I'm livin' at Bandung
  "ITB."
  `
  fmt.Println(msg)

  // empty variable has 'nil' value
  
}
