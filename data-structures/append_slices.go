package main 

// THe append method adds new elements to the slice 
// if the slice capacity has reached the size of the 
// underlying array, then append increases the size 
// by creating a new underlying array and adding the new element. 
// slice1 is a sub slice of arr starting from zero to 3 (excluded) 
// while slice2 is a sub slice of arr strting from 1 (inclusive) to 
// 5 (excluded). In  the following snippet, the append method 
// calls on slice2 to add a new 12 element (append_slice.go) 

import (
	"fmt"
)

func main(){
	var arr = []int{5, 6, 7, 8, 9} 
	var slice1 = arr[:3] 
	fmt.Println("slice1 ", slice1)  
	var slice2 = arr[1:5] 
	fmt.Println("slice2 ", slice2) 
	var slice3 = append(slice2, 12) 
	fmt.Println("Slice 3 ", slice3)
}