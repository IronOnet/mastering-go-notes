package main 

/**
	In addition to lists and maps Go has several more collections 
	available underneath the container package.

	List: 
	the container/list package implements a doubly linked list. 
	A 
**/

import (
	"fmt" 
	"container/list" 
	"sort"
)

func doublyListRoutine(){
	var x list.List 
	x.PushBack(1) 
	x.PushBack(2) 
	x.PushBack(3) 

	for e := x.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value.(int))
	}
}

/**
	Sort: 
	The sort package contains functions for sorting arbitrary data. 
	There are several predefined sorting functions (for slices of ints) 
	and floats. Here's an example for how to sort you own data:
**/ 

type Person struct{
	Name string 
	Age int 
} 

type ByName []Person 

func (this ByName) Len() int{
	return len(this)
}

func (this ByName) Less(i, j int) bool{
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int){
	this[i], this[j] = this[j], this[i]
}

type ByAge []Person 

func (this ByAge) Len() int{
	return len(this)
}

func (this ByAge) Less(i, j int) bool{
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int){
	this[i], this[j] = this[j], this[i]
}

func listSorter(){
	kids := []Person{
		{"Jill", 30}, 
		{"Jack", 33},
	}

	sort.Sort(ByName(kids)) 
	fmt.Println(kids)
}

func main(){
	doublyListRoutine() 
	listSorter()
}