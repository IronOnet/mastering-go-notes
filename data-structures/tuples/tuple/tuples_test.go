package tuple

import (
	"testing" 
)

func TestPowserSeriesNamed(t *testing.T){
	expected_square, expected_cube := 9, 27
	square, cube := PowerSeriesNamed(3) 
	if square != expected_square && cube != expected_cube {
		t.Errorf("Expected Square value of %d and Cube value of %d got %d and %d", expected_square, expected_cube, square, cube)
	}
}

func TestPowserSeries(t *testing.T){
	expected_square, expected_cube := 9, 27
	square, cube := PowerSeries(3) 
	if square != expected_square && cube != expected_cube {
		t.Errorf("Expected Square value of %d and Cube value of %d got %d and %d", expected_square, expected_cube, square, cube)
	}
}
