package main
import "fmt"
func main(){
a := map[string]string{"name":"harika" ,"status":"married","kids" :"two"} //key type string and value type string and intializing directly
b := map[string]int{"abhilash" :32 ,"harika" :30 ,"nayanika" :7,"aadvitha" :4} //key type string and value type int

 var h = make(map[string]string) //empty map
 
h["a"] ="apple"
h["b"] ="ball"
h["c"] ="cat"
h["d"] ="dog"

fmt.Println(h["a"] ,h["b"]) // prints particular key

delete(h,"d") //delets d value in h map
fmt.Println(h)

h["a"] ="ant" //updating an element
fmt.Println(h["a"]) 

 h["e"] = "elephant" //adding element
 h["f"] = "fish"
fmt.Println(h) 

val,ok := a["name"] // checking existance of key and value
fmt.Println(val,ok)

fmt.Println(a) //prints total map
fmt.Println(b)

}