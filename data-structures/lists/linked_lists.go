package main 

import (
	"fmt"
)

type Node struct{
	property int 
	nextNode *Node
}

type LinkedList struct{
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int){
	var node = new(Node) 
	node.property = property  
	node.nextNode = nil 

	if linkedList.headNode != nil{
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = node
}

func (linkedList LinkedList) NodeWithValue(property int) *Node{
	var node *Node 
	var nodeWith *Node 
	for node = linkedList.headNode; node != nil; node = node.nextNode{
		if node.property == property{
			nodeWith = node 
			break
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int){
	var node = new(Node) 
	node.property = property  
	node.nextNode = nil 
	//var nodeWith *Node 
	for currentNode := linkedList.headNode; currentNode != nil; currentNode = currentNode.nextNode{
		if currentNode.property == nodeProperty{
			node.nextNode = currentNode.nextNode
			currentNode.nextNode = node 
			break
		}
	}
}

func (linkedList *LinkedList) LastNode() *Node{
	//var node *Node 
	var lastNode *Node 
	for node := linkedList.headNode; node != nil; node = node.nextNode{
		if node.nextNode == nil{
			lastNode = node 
		}
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int){
	var node = new(Node) 
	node.property = property 
	node.nextNode = nil  

	var  lastNode = linkedList.LastNode() 
	if lastNode != nil{
		lastNode.nextNode = node
	}
}

func (linkedList *LinkedList) IterateList(){
	for node := linkedList.headNode; node != nil; node = node.nextNode{
		fmt.Println(node.property)
	}
}

func main(){
	//var linkedList LinkedList 
	var linkedList = new(LinkedList) 

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(2)
	linkedList.AddToHead(5) 
	linkedList.AddAfter(1, 7) 

	linkedList.IterateList()
}