package main
import "fmt"

func main()  {
  var names [4] string
  names[0] = "andy"
  names[1] = "yandi"
  names[2] = "lukman"
  names[3] = "ardi"

  fmt.Println(names[0], names[1], names[2], names[3])

  // direct declaration
  var fruits = [4] string {"apple","pinneapple","strawberry","orange"}

  for i:=0; i < len(fruits); i++{
    fmt.Printf("element %d : %s \n", i, fruits[i])
  }

  // initialization without prefix value
  var numbers = [...] int{2,3,4,45,45}
  fmt.Println(numbers)

  // multidimension array
  var numbers1 = [2][3]int{[3]int{1,2,3}, [3]int{4,5,6}}
  var numbers2 = [2][2]int{[2]int{1,2}, [2]int{4,5}}

  fmt.Println("numbers1 ", numbers1)
  fmt.Println("numbers2 ", numbers2)

  // for range
  var books = [4]string{"harpot","laspel","dee","doo"}

  for _, book := range books {
    fmt.Printf("nama buku : %s \n", book)
  }

}
