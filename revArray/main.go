package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}

	fmt.Println("Reversed array",reverseArray(arr))
	
	
}

func reverseArray (arr []int) []int{
	x:=len(arr)
	for i:=0;i<x/2;i++{
		arr[i],arr[x-i-1]=arr[x-i-1],arr[i]
	}
	return arr

}
