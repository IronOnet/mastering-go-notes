package main 

// maps are used to keep track of keys that are types, such as integers 
// strings, float, double, pointers, interfaces, structs, and arrays. 
// the values can be of different types. In the following example

import (
	"fmt"
)

func main(){
	var languages = map[int]string{
		3: "English", 
		4: "French", 
		5: "Spanish",
	}
	fmt.Printf("Languages %v", languages)
	
	// Maps can be created using the make method, specifying the key type and the 
	// value type. Products of a map type with a key integer and value string are shown 
	// in the following code 
	var products = make(map[int]string) 
	products[1] = "Chair"
	products[2] = "Table" 
	products[3] = "Stool" 
	products[4] = "Mat"  
	fmt.Printf("The list of Products is %v", products)

}