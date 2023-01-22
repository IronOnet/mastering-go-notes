package main 

/**
In a scenario where class responsabilities are removed or added, the decorator pattern 
is applied. THe decorator pattern helps with subclassing when modifying functionality, 
instead of static inheritance. An object can have multiple decorators and run-time 
decorators. the single responsability principle can be achieved using a 
decorator. 

The Decorator pattern attaches additonal responsibilities to an object 
dynamically. Decorators provide a flexible alternative to subclassing 
for extending functionality.

The decorator pattern participants are the component interface, the concrete component 
class, and the decorator class 

The concrete component implements the component interface 
The decorator class implements the component interface and provides 
additional functionality in the same method or additional methods. 
The decorator base can be a participant representing the base class for 
all decorators. 
*/

import "fmt" 

// IProcess Interface 
type IProcess interface{
	process()
}

// process struct 
type ProcessClass struct{} 

// process class method process 
func (process *ProcessClass) process(){
	fmt.Println("ProcessClass process")
}

// ProcessDecorator struct 
type ProcessDecorator struct{
	processInstance *ProcessClass
}

func (decorator *ProcessDecorator) process(){
	if decorator.processInstance == nil{
		fmt.Println("ProcessDecorator process")
	} else{
		fmt.Println("ProcessDecorator process and ") 
		decorator.processInstance.process()
	}
}

func main(){
	var process = &ProcessClass{} 
	var decorator = &ProcessDecorator{} 
	decorator.process() 
	decorator.processInstance = process 
	decorator.process()
}