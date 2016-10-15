package main

import (
  "fmt"
  "strings"
)
func main(){
  fmt.Println(
    // true
    strings.Contains("test", "st"),

    // 2
    strings.Count("test","t"),

    // true
    strings.HasPrefix("test","te"),

    // true
    strings.HasSuffix("test", "st"),

    // 1
    strings.Index("test","s"),

    // "a-b"
    strings.Join([]string{"a","b"},"-"),

    // == "aaaaa"
    strings.Repeat("a", 6),

    // "bbaaa"
    strings.Replace("aaaa","a","b",2),

    // []strings{"a","b","c","d","e"}
    strings.Split("a-b-c-d-e-", "-"),

    // "test"
    strings.ToLower("TEST"),

    // "TEST"
    strings.ToUpper("test"),
  )

  // convert strings
  arr := []byte("test")
  str := string([]byte{'t','e','s','t'})
  fmt.Println(arr)
  fmt.Println(str)
}
