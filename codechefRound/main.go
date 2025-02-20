package main
import "fmt"

func main(){
	// your code goes here
	var N int
	fmt.Println("Enter the date")
	fmt.Scan(&N)
	if N==4||N==11||N==18||N==25{
	    fmt.Println("YES")
	}else{
	    fmt.Println("NO")
	}
}