package main
import (
  "fmt"
  "net/url"
)

func main(){
  var urlString = "http://google.com/code?name=john wick&age=22"
  var u, e = url.Parse(urlString)

  if e != nil {
    fmt.Println(e.Error())
  }

  fmt.Println("url: %s\n", urlString)
  fmt.Println("protocol: %s\n", u.Scheme) // http
  fmt.Println("host: %s\n", u.Host) // google.com
  fmt.Println("path: %s\n", u.Path) // /code

  var name = u.Query()["name"][0] // json
  var age = u.Query()["age"][0] // standart
  fmt.Println("name: %s, age: %s\n", name, age)
}
