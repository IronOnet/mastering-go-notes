package main 

/**
	The Bridge pattern 
	The bridge pattern decouples the implementation from the 
	abstraction. The abstract base class can be modified 
	easily. THe interface, wich is a bridge, helps in making 
	the functionaity of concrete classes independent from the 
	interface implementer classes.

	The bridge patterns allow the implementation details to 
	change at runtime. this patterns demonstrates the principle 
	"preferring composition over inheritance". It helps 
	in situations where on should subclass multiple time 
	orthogonal to each other. Runtime binding of the app, 
	mapping of orthogonal to each other. 

	Let's say IDrawShape is an interface with the drawshape methjod. 
	DrawShape implements the IDrawShape interface. We create an IContour 
	bridge interface with the drawContour method. The contour class 
	Implements the IContour interface. the ellipese class implements 
	the contour bridge interface to implement the 
	drawContour method. The drawContour method calls the drawShape 
	method on the drawShape instance
**/

import (
	"fmt"
)

// IDraw interface 
type IDrawShape interface{
	drawShape(x[5] float32, y[5] float32)
}

// DrawShape struct 
type DrawShape struct{}

// THe drawshape method draws the shape given the 
// coordinates. as shown in the following code 

// DrawShape struct has method draw Shape with float x and y coordinates 
func (drawShape DrawShape) drawShape(x[5] float32, y[5] float32){
	fmt.Println("Drawing Shape")
}

// IContour interface 
type IContour interface{
	drawContour (x[5] float32, y[5] float32) 
	resizeByFactor(factor int)
}

// DrawContour struct 
type DrawContour struct{
	x[5] float32 
	y[5] float32 
	shape DrawShape 
	factor int 
}

// THe drawContour method of the DrawContour class calls the drawShape method 
// on the shape instance, this is shown in the following code: 

// DrawContour method drawContour given the coordinates 
func (contour DrawContour) drawContour(x[5] float32, y[5] float32){
	fmt.Println("Drawing Contour") 
	contour.shape.drawShape(contour.x, contour.y)
}

// DrawContour method resizedByFactor given factor 
func (contour *DrawContour) resizeByFactor(factor int){
	contour.factor = factor 
}

// main method 
func main(){
	var x = [5]float32{1, 2, 3, 4, 5} 
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = &DrawContour{x, y, DrawShape{}, 2} 
	contour.drawContour(x, y) 
	contour.resizeByFactor(2)
}