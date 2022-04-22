package main
import "fmt"
func main(){
var large , new int
x :=[5]int{3,4,9,6,2} // finding the largest number

for _,value := range x{
if value > new{
new = value
large = new
}
} 
fmt.Println("largest num is ", large)

}