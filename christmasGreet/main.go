package main
import "fmt"

func main(){
    var date int
	fmt.Println("Enter the date of december")
	fmt.Scan(&date)
	if (date<1 || date >31){
	    fmt.Println("Wrong input is given")
	}else if date==25{
	fmt.Println("CHRISTMAS")
	    
	}else{
	    fmt.Println("ORDINARY")
	}
	
	
}

// package main

// import "fmt"

// func main() {
//     var X int
//     fmt.Scan(&X)  // Read the input

//     if X == 25 {
//         fmt.Println("CHRISTMAS")
//     } else {
//         fmt.Println("ORDINARY")
//     }
// }
