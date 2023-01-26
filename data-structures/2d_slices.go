package main  

import (
	"fmt"
)



func main(){
	fmt.Println("2D slices")
	var TowDArray [8][8]int 
	TowDArray[3][6] = 18 
	TowDArray[7][4] = 3 
	fmt.Println(TowDArray) 

	// For dynamic allocation, we use slice of slices. In the following code, slice of slice 
	// is explained as two-dimentionsl slices -
	var rows, cols int  
	rows, cols  = 7, 9 
	var twodslices = make([][]int, rows) 
	for i, _ := range twodslices{
		twodslices[i] = make([]int, cols)
	}

	fmt.Println(twodslices)
}