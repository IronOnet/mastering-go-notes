package main 

import (
	"fmt" 
	"math"
)

func distance(x1, y1, x2, y2 float64) float64{
	a := x2 - x1 
	b := y2 - y1 
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64{
	l := distance(x1, y1, x1, y2) 
	w := distance(x1, y1, x2, y2) 
	return l * w 
}

func circleArea(x, y, r float64) float64{
	return math.Pi * r * r
}

/** Structs 

An easy way to make this program better is to use a struct. 
A struct is a type which contains named fields 
For example we could represent a circle like this 
**/

type Circle struct{
	x, y, z float64 
	
}

func (c *Circle) area() float64{
	return math.Pi * c.z * c.z 
}

type Rectangle struct{
	x1, y1, x2, y2 float64 

}


func (r *Rectangle) area() float64{
	l := distance(r.x1, r.y1, r.x1, r.y2) 
	w := distance(r.x1, r.y1, r.x2, r.y1) 
	return l * w 
}


/** 
	Embedded types 
	A struct's field usually represent the has-a relationship. for example a 
	Circle has a radius. Suppose we had a person struct: 

	type Person Struct{
		Name string 
	}

	func (p *Person) Talk(){
		fmt.Println("Hi, my name is", p.Name)
	}

	And we wanted to create a new Android struct. We 
	could do this 
	type Android struct{
		Person Person 
		Model string 
	}

	This would work , but we would rather say an android is a Person rather 
	than an ANdroid has a person. Go supports relationships like this 
	by using an embedded type. Also known as anonymous fields 

*/ 

type Person struct{
	Name string 
}

func (p *Person) Talk(){
	fmt.Println("Hello my name is: ", p.Name)
}

type Android struct{
	Person 
	Model string 
}



func typeExample(){
	// Type initialization 
	// we can create an instance of our new Circle type 
	// in a variety of ways 
	var c Circle 
	// Like with other datatypes this, will crete a local Circle 
	// variable that is by default set to zero. for a struct 
	// zero means each of the fiels is set to their corresponding 
	// zero value (0 for ints, 0.0 for flaots) 
	// nil for pointers, we can also use the new function
	c2 := new(Circle) // this allocates memory for all the fields, set 
	// each of them to their zero value and returns a pointer. 
	c = Circle{x: 0, y: 0, z: 5} 
	c2 = &Circle{x: 0, y: 0, z: 5}
	fmt.Println(c, c2) 

}

/**
	Interfaces: 
	You may have noticed that
	Like a struct an interface is created using the type keyword, followed by a 
	name and the keyword interface. BUt instead of defnining fields, 
	we define a "method set". A method set is a list of method that a 
	type must have in order to "implement" the interface.
**/

func totalArea(shapes ...Shape) float64{
	var area float64 
	for _, s := range shapes{
		area += s.area()
	}
	return area 
}

type Shape interface{
	area() float64
}

type LogicProvider struct{} 

func (lp LogicProvider) Process(data string) string{
	return "Hello World " + data 
}

type Logic interface{
	Process(data string) string 
}

type Client struct{
	L Logic 
}

func (c Client) Program(data string){
	c.L.Process(data)
}

func main(){
	var rx1, ry1 float64 = 0, 0 
	var rx2, ry2 float64 = 10, 10 
	var cx, cy, cr float64 = 0, 0, 5 
	// r := new(Rectangle) 
	// c := new(Circle) 
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2)) 
	fmt.Println(circleArea(cx, cy, cr))
//	fmt.Println(totalArea(&c, &r))
}