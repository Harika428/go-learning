package main
import ("fmt")

type circle struct {
 x,y,z int
}
func circleArea(c *Circle) float64 {
return math.Pi * c.x*c.x
}
 func main(){
 var c circle
 c.x=3
 c.y=22
 c.y = 33
 fmt.Println(c.x,c.z)
 fmt.Println(circleArea(c))
 circleArea()
 }