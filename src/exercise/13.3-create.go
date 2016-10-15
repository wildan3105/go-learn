package main
import "os"

func main(){
  file, err := os.Create("file.txt")
  if err != nil {
    // handle error here
    return
  }
  defer file.Close()
  file.WriteString("hello, this is a dummy file")
}
