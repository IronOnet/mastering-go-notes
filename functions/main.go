package main 

import (
	"fmt"
)

// Returning multiple values with functions 

func f()(int, int){
	return 5, 6
}

// 

// Variadic functions 
func add(args ...int) int{
	total := 0 
	for _, v := range args{
		total += v 
	}
	return total 
}


// Closures 
// it is possible to create functions inside of functions 
func aClosure(){
	a := func() int {
		return 119
	}

	add := func(x, y int) int{
		return x + y
	}

	getAhundredNineteen := a() 
	resultSum := add(10, 20)
	fmt.Println(getAhundredNineteen, resultSum)
}

func makeEvenGenerator() func() uint{
	// Unlike nomral variables this local 
	// variable persists between calls.
	i := uint(0) 
	return func() (ret uint){
		ret = i 
		i += 2  
		return 
	}
}



func main(){
	v, k := f() 
	fmt.Printf("\nValues returned by f(): %d, %d", v, k)
	fmt.Println("\nVariadic add result for (1, 2, 3)", add(1, 2, 3))

	// 
	fmt.Println("\n********Closures**********")
	nextEven := makeEvenGenerator() 
	fmt.Println(nextEven()) // 0 
	fmt.Println(nextEven()) //2 
	fmt.Println(nextEven()) 
	fmt.Println(nextEven())

}

