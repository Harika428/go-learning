package main
import "fmt"
func main(){
var x [5]int
   x[0]= 44
   x[1]=55
   x[2]=66
   x[3]=77
   var total = 0
for i :=0; i<=4; i++{

 total += x[i]
 }
   fmt.Println(x)
   fmt.Println("total =", total)
}