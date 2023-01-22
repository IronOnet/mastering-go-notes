package main

import (
	"fmt"
	"tuples/tuple"
)

// gets the power series of integer a and returns tuple of square of a and
// a cube of a




// the main method calls the powerSeries method with 3 as parameter
// the square, cubes values are returned from the method 
func main(){
	var square int 
	var cube int 
	square, cube = tuple.PowerSeries(3)

	fmt.Println("Square ", square, " Cube ", cube) 
}
