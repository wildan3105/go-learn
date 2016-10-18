package main
import (
  "encoding/json"
  "fmt"
)
type User struct {
  FullName string `json:"Name"`
  Age int
}

func main(){
  var jsonString = `{"Name": "john wick", "Age": 22}`
  var jsonData = []byte(jsonString)

  var data User

  var err = json.Unmarshal(jsonData, &data)
  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Println("User : ", data.FullName)
  fmt.Println("Age : ", data.Age)
}
