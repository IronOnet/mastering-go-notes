package main 

import (
	"fmt"
)

type Set struct{
	integerMap map[int]bool 
}

func (set *Set) New(){
	set.integerMap = make(map[int]bool)
}

func (set *Set) AddElement(element int){
	if !set.ContainsElement(element){
		set.integerMap[element] = true 
	}


}

func (set *Set) DeleteElement(element int){
	delete(set.integerMap, element) 
}

func (set *Set) ContainsElement(element int) bool{
	var exists bool 
	_, exists = set.integerMap[element]

	return exists
}

func (set *Set) Intersect(anotherSet *Set) *Set{
	var intersectSet = &Set{} 
	intersectSet.New()
	for value , _ := range set.integerMap{
		if anotherSet.ContainsElement(value){
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set{
	var unionSet = &Set{} 
	unionSet.New() 
	for value, _ := range set.integerMap{
		unionSet.AddElement(value)
	}

	for value, _:= range anotherSet.integerMap{
		unionSet.AddElement(value)
	}

	return unionSet
}

func main(){
	// object's instanciation
	set := new(Set) 
	// creates an empty set
	set.New() 

	set.AddElement(1) 
	set.AddElement(2)

	fmt.Println("initial set", set) 
	fmt.Println(set.ContainsElement(1)) 

	anotherSet := new(Set) 
	anotherSet.New() 

	anotherSet.AddElement(2) 
	anotherSet.AddElement(4) 
	anotherSet.AddElement(5) 

	fmt.Println("another set", set) 
	fmt.Println("intersection of sets ", set.Intersect(anotherSet)) 

	fmt.Println("union of sets", set.Union(anotherSet))
}