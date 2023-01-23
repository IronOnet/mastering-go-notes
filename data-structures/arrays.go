package main 

import "fmt"
// Arrays are the most famous data structures in different programming 
// languages. Different data types can be handled as elements in arrays 
// such as int, float32, double and others 

//var arr = [5]int{1, 2, 4, 5, 6}  

// An array's size can be found with the len() function. A for loop is used for 
// accessin all the elements in an array as follows 
// for i:= 0; i < len(arr) ; i++{
// 	fmt.Println("printing elements", arr[i])
// }

func main(){
	var arr = [5]int{1, 2, 4, 5, 6} 
	for i:= 0; i < len(arr); i++{
		fmt.Println("printing elements", arr[i])
	}

	// We can use the range operator to retrive indexes and values 
	for _, value := range arr{
		fmt.Println(" range ", value)
	}
}