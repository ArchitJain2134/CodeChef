package main
import "fmt"

func main(){
	var A,X,Y int
	fmt.Println("Enter the cost of socks you want to purchase")
	fmt.Scan(&A)
	fmt.Println("Enter the Total money saved")
	fmt.Scan(&X)
	fmt.Println("Enter the total money recieved by parents")
	fmt.Scan(&Y)
	
	if (X+Y>=A) {
	    fmt.Println("YES")
	}else{
	    fmt.Println("NO")
	}
	
}