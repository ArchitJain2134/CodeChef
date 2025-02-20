package main
import "fmt"

func main(){
	// your code goes here
	var X int
	fmt.Println("Enter the date")
	fmt.Scan(&X)
	if X>25{
	    fmt.Println("Christmas already gone")
	}else{
	    n:=25-X
	    fmt.Println(n)
	}
}