package main 


// Go has a built-in type for errors that was already seen (the error type) 
// we can cxreate our own errors by using the new function in the 
// errors package 
import "errors"
import "fmt"

func main(){
	err := errors.New("Error Message")
	fmt.Println(err)
}
