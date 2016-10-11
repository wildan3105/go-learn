package main
import "fmt"

func main()  {
  // slice, use built-in function
  x := make([]int, 5)
  i := 0;
  for i<5{
    x[i]=i
    i++
    fmt.Println(i)
  }

  // append
  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)

  // copy
  copy1 := []int{1,2,3}
  copy2 := make([]int, 2)
  copy(copy2, copy1)
  fmt.Println(copy1, copy2)

  // map
  fmt.Println("maps : ")
  y := make(map[int]int)
  y[1] = 10
  fmt.Println(y[1])
  delete(y,1)
  fmt.Println(y[1])

  // elements := make(map[string]string)
  // elements["H"] = "Hydrogen"
  // elements["He"] = "Helium"
  // elements["C"] = "Carbon"
  // name, ok := elements["Un"]
  // fmt.Println(name, ok)
  // fmt.Println(elements["C"])
  // MAPS with JSON

  elements := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
  }
  fmt.Println(elements["He"]["name"])
}
