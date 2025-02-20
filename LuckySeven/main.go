package main
import "fmt"

func main(){
	// your code goes here
	var S string
	fmt.Println("Enter your string")
	fmt.Scan(&S)
	b:=string(S[6])
	fmt.Println(b)
}