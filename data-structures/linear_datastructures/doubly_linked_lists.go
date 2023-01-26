package main  

import (
	"fmt"
)

type Node struct{
	property int 
	nextNode *Node 
	previousNode *Node 
}

type LinkedList struct{
	headNode *Node
}

func (list *LinkedList) AddToHead(property int){
	node := new(Node) 
	node.property = property  
	node.nextNode = nil  
	if list.headNode != nil{
		node.nextNode = list.headNode 
		list.headNode.previousNode = node 
	}

	list.headNode = node 
}

func (list *LinkedList) NodeWithValue(property int) *Node{
	nodeWith := new(Node) 
	for node := list.headNode; node != nil; node = node.nextNode{
		if node.property == property{
			nodeWith = node 
			break
		}
	}

	return nodeWith
}

func (list *LinkedList) AddAfter(nodeProperty int, property int){
	node := new(Node) 
	node.property = property 
	node.nextNode = nil  

	nodeWith := list.NodeWithValue(nodeProperty) 
	if nodeWith != nil{
		node.nextNode = nodeWith.nextNode 
		node.previousNode = nodeWith  
		nodeWith.nextNode = node 
	}
}

func (list *LinkedList) LastNode() *Node{
	lastNode := new(Node) 
	for node := list.headNode; node != nil; node = node.nextNode{
		if node.nextNode == nil{
			lastNode = node 
		}
	}

	return lastNode 
}

func (list *LinkedList) AddToEnd(property int){
	node := new(Node) 
	node.property = property  
	node.nextNode = nil  

	lastNode := list.LastNode() 

	if lastNode != nil{
		lastNode.nextNode = node  
		node.previousNode = lastNode 
	}
}

func (list *LinkedList) IterateList(){
	for node := list.headNode; node != nil; node = node.nextNode{
		fmt.Println(node.property)
	}
}

func (list *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node{
	nodeWith := new(Node) 
	for node := list.headNode; node != nil; node = node.nextNode{
		if node.previousNode != nil && node.nextNode != nil{
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty{
				nodeWith = node 
				break
			}
		}

	}
	return nodeWith
}

func main(){
	linkedList := new(LinkedList) 
	linkedList.AddToHead(1) 
	linkedList.AddToHead(3) 
	linkedList.AddToEnd(5) 
	linkedList.AddAfter(1, 7) 
	fmt.Println(linkedList.headNode.property) 

	node := linkedList.NodeBetweenValues(1, 5) 
	fmt.Println(node.property)
}