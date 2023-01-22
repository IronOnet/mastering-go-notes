package main 

/**
	Composite design pattern 
	A composite is a group of similar objects in a single object. Objects 
	are stored in a tree form to persist the whole hierarchy. 
	THe composite pattern is used to change a hierarchical collections 
	of objects. THe compossite pattern is modeled on an heterogenous 
	collection.  

	The composite patter comprises the compoenet interface, component class, 
	composite and client. 


	* The component interface defines the default behavior of all objects 
	and behaviors for accessing the components of the composite. 

	* The composite and component classes implement the component interface. 
	* THe client interacts with the component interface to invoke methods in 
	the composite.
*/

import "fmt" 

// IComposite interface 
type IComposite interface{
	perform()
}

// Leaflet struct 
type Leaflet struct{
	name string 
}

// Leaflet class method perform 
func (leaf *Leaflet) perform(){
	fmt.Println("Leaflet " + leaf.name)
}

// Brach struct 
type Branch struct{
	leafs []Leaflet 
	name string 
	branches []Branch
}

// THe perform method of the branch class calls the perform method on 
// branch and leafs 

// Branch class method perform 
func (branch *Branch) perform(){
	fmt.Println("Branch: " + branch.name) 
	for _, leaf := range branch.leafs{
		leaf.perform()
	}

	for _, branch := range branch.branches{
		branch.perform()
	}
}

// Branch class method add leaflet 
func (branch *Branch) add(leaf Leaflet){
	branch.leafs = append(branch.leafs, leaf)
}

func (branch *Branch) addBranch(newBranch Branch){
	branch.branches = append(branch.branches, newBranch)
}

// Branch class method getLeaflets 
func (branch *Branch) getLeaflets() []Leaflet{
	return branch.leafs
}

func main(){
	var branch = &Branch{name: "branch 1"} 
	var leaf1 = Leaflet{name : "leaf 1"} 
	var leaf2 = Leaflet{name: "leaf 2"} 
	var branch2 = Branch{name: "branch 2"} 
	branch.add(leaf1) 
	branch.add(leaf2) 
	branch.addBranch(branch2) 
	branch.perform()
}