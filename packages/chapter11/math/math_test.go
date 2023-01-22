package math 

import "testing" 

type testPair struct{
	values []float64 
	average float64
}

type testPairInt struct{
	values []int 
	result int
}

/**
	This is a common way to setup tests. We create a struct 
	to represent the inputs and outputs for the functon. then 
	we create a list of these struct (pairs). THen we loop 
	through each one and run the function.
**/

var tests = []testPair{
	{ []float64{1, 2}, 1.5}, 
	{ []float64{1, 1, 1, 1, 1, 1}, 1}, 
	{ []float64{-1, 1}, 0},
}

var testsMin = []testPairInt{
	{[]int{1, 2}, 1}, 
	{[]int{3, 2}, 2}, 
	{[]int{8, 32}, 8},
}

func TestMin(t *testing.T){
	for _, pair := range testsMin{
		r := Min(pair.values[0], pair.values[1])
		if r != pair.result{
			t.Errorf("For Min(%d, %d) -----Got :%d -----Expected: %d", pair.values[0], pair.values[0],
		r, pair.result)
		}
	}
}

var testsMax = []testPairInt{
	{[]int{1, 2}, 2}, 
	{[]int{3, 2}, 3}, 
	{[]int{8, 32}, 32},
}


func TestMax(t *testing.T){
	for _, pair := range testsMax{
		r := Max(pair.values[0], pair.values[1]) 
		if r != pair.result{
			t.Errorf("For Min(%d, %d) -----Got :%d -----Expected: %d", pair.values[0], pair.values[0],
		r, pair.result)
		}
	}
}

func TestAverage(t *testing.T){
	for _, pair := range tests{
		v := Average(pair.values) 
		if v != pair.average{
			t.Error(
				"For", pair.values, 
				"expected", pair.average, 
				"got", v,
			)
		}
	}
}