package main 

import (
	"container/heap" 
	"fmt"
)

type IntegerHeap []int 

func (iheap IntegerHeap) Len() int { return len(iheap)} 

// IntegerHeap method - checks if element of i index is less than j index 
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j]} 

// IntegerHeap method - swaps the element at index of i to j index
func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i]}  

// IntegerHeap has a Push method that pushes the item with the interface 

// IntegerHeap method - pushes the item 
func (iheap *IntegerHeap) Push(heapintf interface{}){
	*iheap = append(*iheap, heapintf.(int))
}

// IntegerHeap method - pops the time from the heap 
func (iheap *IntegerHeap) Pop() interface{}{
	var n int 
	var x1 int  
	var previous IntegerHeap = *iheap 
	n = len(previous) 
	x1 = previous[n-1] 
	*iheap = previous[0:n-1] 
	return x1
}

// main method 
func main(){
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5} 
	heap.Init(intHeap)
	heap.Push(intHeap, 2) 
	fmt.Printf("minimum: %d\n", (*intHeap)[0]) 
	for intHeap.Len() > 0{
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}

