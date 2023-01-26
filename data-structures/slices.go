// Slices. Go Slice is an abstration over Go Array. Multple data eelemnts 
// of the same type are allowed by Go Arrays. The definition of variables 
// that can hold several data elements of the same type are allowed 
// by Go array, but it does not have hany provision of inbuilt 
// methods to increaset its size. 
package main 

import (
	"fmt"
)



// The slice function. slices are passed by referring to functions. 
// Big slices can be passed to functions without impacting performance. 
// Passing a slice as a reference to a function is demonstrated 

// twiceValue method given slice of int type 
func twiceValue(slice []int){
	var i int 
	var value int  

	for i, value = range slice{
		slice[i] = 2*value
	}
}


func main(){
	var slice  = []int{1, 3, 5, 6} 
	slice = append(slice, 8) 
	fmt.Println("Capacity", cap(slice)) 
	fmt.Println("Length", len(slice))  

	var sliceII = []int{1, 3, 5, 6} 
	twiceValue(sliceII) 
	for i := 0; i < len(sliceII); i++{
		fmt.Println("new slice value", slice[i])
	}
	fmt.Println("The final slic: ", sliceII)


	
}