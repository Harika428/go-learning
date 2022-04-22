package main
import "fmt"

func main(){
var small int  
x  := [6]int{56, 22,44,55,99,33} // finding the smallest number
small = x[0]
for _,value := range x { 
 if value < small {
 small = value
  }
} 

fmt.Println("smallest num is", small)

}