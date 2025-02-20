package main
import "fmt"

func main(){
	// your code goes here
	var X,Y int
	fmt.Println("Enter the number of cones")
	fmt.Scan(&X)
	fmt.Println("Enter the number of ice cream")
	fmt.Scan(&Y)
	if X==Y{
	    fmt.Println(X)
	}else if X<Y{
	    fmt.Println(X)
	}else {
	    fmt.Println(Y)
	}
}