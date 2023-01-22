package tuple

func PowerSeries(a int)(int, int){
	return a*a, a*a*a 
}

// power series implementation with named returns
func PowerSeriesNamed(a int) (square int, cube int){
	square = a * a 
	cube = square * a 
	return
}