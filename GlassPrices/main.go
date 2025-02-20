package main
import "fmt"

func main(){
	// your code goes here
	var X,Y int
	fmt.Println("Enter the cost of plastic frame")
	fmt.Scan(&X)
	fmt.Println("Enter the cost of metallic frame")
	fmt.Scan(&Y)
	if Y>2*X{
	    fmt.Println("PLASTIC")
	}else{
	    fmt.Println("METAL")
	}
}
