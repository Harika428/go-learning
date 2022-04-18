package main
import "fmt"
func main(){
arr :=[3]int{2,3,4}
arr1 :=[6]int{1,2,3,4,5}
arr2 :=[]string{"i","have","a","book"}
arr4 :=[6]int{2:4,5:5} //partially intialised
  
 myslice := arr1[2:4] // slice created from array
 
 fmt.Println(myslice)
 
// this prints 2nd element//
fmt.Println(arr[1])


// changes elements of array
arr1[5] = 7
arr1[4] = 6
fmt.Println(arr1)

fmt.Println("length =", len(arr2)) // this prints length of an array
fmt.Println("capacity =",cap(arr)) //prints capacity of an array

//this prints total string
fmt.Println(arr2)
fmt.Println(arr4)
}