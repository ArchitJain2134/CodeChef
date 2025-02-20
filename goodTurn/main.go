package main

import "fmt"

func main() {
    var t int
	fmt.Println("Enter the number of test")

    fmt.Scanf("%d", &t)

    // Loop for each test case
    for t > 0 {
        var x, y int
        fmt.Println("Enter the no. on dice of chef and chefina resptively")
        fmt.Scanf("%d %d", &x, &y)

        // Your code for each test case goes here
        if x+y>6{
            fmt.Println("YES")
        }else{
            fmt.Println("NO")
        }

        t--
    }
}