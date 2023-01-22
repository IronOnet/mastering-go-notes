package main 

/**
	When we call a function that takes an argument, that 
	argument is copied to the function:
**/
import (
	"fmt"
)

func zero(x int){
	x = 0 
}

func zeroWPtr(xPtr *int){
	*xPtr = 0
}

// Pointers reference a location in memory where a value 
// is stored rather than the value itself. (They point to 
// something else) By using a pointer (*int) the zero 
// function is able to modify the original variable.

/** 
The * and & operators 
	In Go a pointer is represented using the * asterisk characther 
	followed by the ype of the stored value. In the zero fuction 
	xPtr is a pointer to an int.

	* is also used to "dereference" pointer variables. Dereferencing a pointer
	gives us access to the value the pointer points to. When we write 
	*xPtr = 0, we are saying "store the int 0 in the memory location xPtr referes to" 
	If we try xPtr = 0 instead we will get a compiler error because xPtr is not an int
	it's a *int, which can only be given another *int 

	Finally we use the & operator to find the address of a variable. &x return a *int (pointer to an int) 
	because x is an int. This what allows us to modify the original variable. &x in main 
	xPtr in zero refer to the same memory locatio


new: 

Another way to get a pointer is to use the built-in new function 
**/

func one(xPtr *int){
	*xPtr = 1
}



func main(){
	x := 5 
	zero(x) 
	fmt.Println(x) // X is still 5 
	zeroWPtr(&x) // x is now 0 
	fmt.Println(x) 

	xPtr := new(int) 
	one(xPtr) 
	fmt.Println(*xPtr)
}