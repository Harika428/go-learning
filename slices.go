package main

import "fmt"

var slice[]int

func main() {
 slice1:= []int{2,3,4,5,6,7}
 fmt.Println(slice1)

 slice2:=append(slice1,0,1)
fmt.Println(slice2)

slice:= slice1[2:6]
  fmt.Println(slice) 
  
slice3:=make([]int,3)
copy(slice3,slice2)
fmt.Println(slice3)
}
  
