package main

import "fmt"

func main() {
	// your code goes here
	var X int
	fmt.Println("Enter the date")
	fmt.Scan(&X)
	if X == 1 || X == 2 || X == 3 || X == 4 || X == 5 || X == 6 || X == 7 || X == 8 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
