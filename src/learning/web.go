package main
import (
  "net/http"
  "fmt"
  "html/template"
)

func	main()	{
				http.HandleFunc("/",	func(w	http.ResponseWriter,	r	*http.Request)	{
								var	data	=	map[string]string{
												"Name":	"john	wick",
												"Message":	"have	a	nice	day",
								}
								var	t,	err	=	template.ParseFiles("template.html")
								if	err	!=	nil	{
												fmt.Println(err.Error())
								}
								t.Execute(w,	data)
				})
				fmt.Println("starting	web	server	at	http://localhost:9000/")
				http.ListenAndServe(":9000",	nil)
}
