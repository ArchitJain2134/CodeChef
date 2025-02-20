package main
import "fmt"

func main(){
	// your code goes here 
	var X , Y ,Z int
	fmt.Println("Enter the no of Rainy days in a week")
	fmt.Scan(&X)
	fmt.Println("Enter the no of cloudy days in a week")
	fmt.Scan(&Y)
	Z=7-(X+Y)
	fmt.Println(Z)
}