package main 

/**
FLyweight is used to manage the state of an object with high variation. THe pattern 
allows us to share common parts of the object state among multiple objects, 
instead of each object storing it. Variable object data is reffered to as extrinsic 
state, and the rest of the object state is intrinsic.  

FLyweight objects are immutable. Value objects are good example of the flyweight pattern. 
Flyweight objects can be created in a single thread mode, ensuring one instance per value. 
In a concurrent thread scenario, multiple instances are created. THis is based on the 
equality criterion of flyweight objects. 

The participatns of the flyweight pattern are the Flyweight interface 
ConcreteFlyWeight, FlyweightFactory, and the Client classes
*/ 

import (
	"fmt" 
)

// DataTransferObjectFactory struct 
type DataTransferObjectFactory struct{
	pool map[string] DataTransferObject
}

// DataTransferObjectFactory class method getDataTransferObject 
func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject{
	var dto = factory.pool[dtoType] 
	if dto == nil{
		fmt.Println("new DTO of dtoType: " + dtoType) 
		switch dtoType{
		case "customer": 
			factory.pool[dtoType] = Customer{id: "1"} 
		case "employee": 
			factory.pool[dtoType] = Employee{id: "2"} 

		case "manager": 
			factory.pool[dtoType] = Manager{id:"3"} 
		case "address": 
			factory.pool[dtoType] = Address{id: "4"}
		}
	
		dto = factory.pool[dtoType]
	}

	return dto
}


// DataTransferObject interface 
type DataTransferObject interface{
	getId() string
}

// Customer struct 
type Customer struct{
	id string
	name string 
	ssn string
}

// Customer class method getId 
func (customer Customer) getId() string{
	return customer.id 
}

// Employee struct 
type Employee struct{
	id string 
	name string
}

// Employee class method getId 
func (employee Employee) getId() string{
	return employee.id
}

// Manager struct 
type Manager struct{
	id string 
	name string 
	dept string 
}

// The DataTransferObject interface is implemented by the Manager class, as shown 
// in the following code 

// Manager class method getId 
func (manager Manager) getId() string{
	return manager.id 
}

// Address struct 
type Address struct{
	id string 
	streeLine1 string 
	streetLine2 string 
	state string 
	city string
}

// Adress class method getId 
func (address Address) getId() string{
	return address.id
}

func main(){
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)} 
	var customer DataTransferObject = factory.getDataTransferObject("customer") 
	fmt.Println("Customer ", customer.getId()) 
	var employee DataTransferObject = factory.getDataTransferObject("employee") 
	fmt.Println("Employee ", employee.getId()) 
	var manager DataTransferObject = factory.getDataTransferObject("manager") 
	fmt.Println("Manager", manager.getId()) 
	var address DataTransferObject = factory.getDataTransferObject("address") 
	fmt.Println("Address", address.getId())
}
