package main
import "fmt"

func main()  {
  // variable declaration
  var firstName string = "Wildan"
  var lastName = "Syahrun"
  middleName := "Nahar"

  // multiple var declaration
  var first, second, third string
  first, second, third = "satu", "dua", "tiga"

  // multiple types of vars
  nama, umur, lahir, cowok := "WSN", 20, 1994, true
  var freeVar = "freelancer"

  // trash var
  _ = "learning GO"

  fmt.Printf("BIODATA : %s %s %s %s", nama, umur, lahir, cowok, freeVar)
  fmt.Printf("This is an example : %s %s %s \n", first, second, third)

  fmt.Printf("halo Wildan Syahrun! \n")
  fmt.Printf("halo %s %s %s!\n", firstName, lastName, middleName)
  fmt.Println("halo", firstName, lastName + "!")
}
