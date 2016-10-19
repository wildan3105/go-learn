package main
import (
  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main()  {
  router := gin.Default()
  router.LoadHTMLGLob("templates/*")

  // default route
  router.GET("/", func(c *gin.Context){
    // call the THML method of the Context to render a template
    c.HTML(
      http.StatusOK,
      "index.html",
      gin.H{
        "title":"Home Page",
      },
    )
  })

  // start app
  router.Run()
}
