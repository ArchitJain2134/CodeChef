package main
import "fmt"

func main(){
	// your code goes here
	var N int
	fmt.Println("Enter the date of march")
	fmt.Scan(&N)
	if (N<1 || N>31){
	    fmt.Println("Wrong input is given")
	}else if (N==15){
	    fmt.Println("yes")
	}else {
	    fmt.Println("NO")
	}
	
}