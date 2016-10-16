package main
import (
  "net/http"
  "io"
)

func hello(res http.ResponseWriter, req *http.Request){
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
    res,
      `
        <doctype html>
          <html>
            <head>
              <title> Go Web! </title>
            </head>
            <body>
              <h1> Hi, this is my 1st Go Web :) </h1>
            </body>
          </html>
      `,
  )
}
func main(){
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":9000", nil)
}
