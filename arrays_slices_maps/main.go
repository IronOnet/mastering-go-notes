package main 

import (
	"fmt"
)

func makeArrays(){
	 arr := []int{1, 2, 3} 
	 arrLen := len(arr) 
	 fmt.Printf("The length of the array %d is: %d", arr, arrLen)
}

// An array is a numbered sequence of elements of a single type 
// with a fixed length. In Go they look like this/ 
func someArray() (someArray []int){
	someArray = make([]int, 5) 
	someArray[0], someArray[1], someArray[3] = 3, 2, 1
	return 
}

func floatArrays(){
	var x [5]float64 
	x[0] = 79 
	x[1] = 94 
	x[2] = 92 
	x[3] = 88 
	x[4] = 2023 

	total := float64(0) 
	for i:= 0; i < len(x); i++{
		total += x[i]
	}
	fmt.Println(total/5)
}

/** Maps 
	A map is an unordered collection of key-value pairs. Also known as 
	an associative array, a hash table or a dictionary, maps are used 
	to look up a value by its associated keys.  
**/ 

func mapsExample(){
	var x map[string] int 
	x  = make(map[string]int)
	x["Some String"] = 13

	y := make(map[string]int) 
	y["key"] = 10 
	fmt.Println(y["key"])
}

func computeXS(series []float64) float64{
	//xs := []float64{98, 93, 77, 82, 83} 
	total := 0.0 
	for _, v := range series{
		total += v 
	}
	result := total/ float64(len(series)) 
	return result
}


func main(){
	makeArrays() 
	arr := someArray() 
	xs := []float64{98, 93, 77, 82, 83}
	result := computeXS(xs) 
	fmt.Printf("\nAverage of series %f is: %f", xs, result)
	fmt.Println("\n",arr)
	floatArrays()
	mapsExample()
}