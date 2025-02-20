package main
import "fmt"

func main(){
	// your code goes here
	var x,y int 
	fmt.Println("Enter the number of donuts chef ate")
	fmt.Scan(&x)
	fmt.Println("Enter the number of calories each dout have")
	fmt.Scan(&y)
	s:=x*y
	fmt.Println(s) // shows total calories
}