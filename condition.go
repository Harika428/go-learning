package main
import "fmt"
func main(){
// using for loop and if conditions//
 for i :=1; i <= 10; i++ {
	if i%2 == 0 {
      fmt.Println("even")
    } else {
	  fmt.Println("odd")
	}
	fmt.Println(i)
  }
}


