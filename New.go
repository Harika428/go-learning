package main
import ("fmt")

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  var pers1 Person
  pers1.name = "harika"
  pers1.age = 30
  pers1.job = "housewife"
  pers1.salary =1000
  fmt.Println(pers1)
  }