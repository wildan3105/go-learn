package main
import (
  "fmt"
  "database/sql"
  "os"
  _ "github.com/go-sql-driver/mysql"
)

type student struct {
  id string
  name string
  age int
  grade int
}

func connect() *sql.DB {
  var db, err = sql.Open("mysql", "root:wildansnahar@tcp(127.0.0.1:8080)/go")
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(0)
  }
  return db
}

func sqlQuery(){
  var db = connect()
  defer db.Close()

  var age = 27
  var rows, err = db.Query("select id, name, grade from tb_student where age = ?", age)

  if err != nil {
    fmt.Println(err.Error())
    os.Exit(0)
  }
  defer rows.Close()

  var result []student

  for rows.Next(){
    var each = student{}
    var err = rows.Scan(&each.id, &each.name, &each.grade)

    if err != nil {
      fmt.Println(err.Error())
      os.Exit(0)
    }
    result = append(result, each)
  }

  if err = rows.Err(); err != nil {
    fmt.Println(err.Error())
    os.Exit(0)
  }
  for _, each := range result {
    fmt.Println(each.name)
  }
}

func main(){
  sqlQuery()
}
