package main 

import (
	"fmt"
)

type Queue []*Order

type Order struct{
	priority int 
	quantity int 
	product string 
	customerName string 
}

func (order *Order) NewOrder(priority int, quantity int, product string, customerName string){
	order.priority = priority 
	order.quantity = quantity  
	order.product = product 
	order.customerName = customerName 


}

func (queue *Queue) Add(order *Order){
	if(len(*queue)) == 0{
		*queue = append(*queue, order)
	} else{
		appended := false 
		for i, addedOrder := range *queue{
			if order.priority > addedOrder.priority{
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...) 
				appended = true 
				break 
			}
		}
		if !appended{
			*queue = append(*queue, order)
		}
	}
}

func main(){
	queue := make(Queue, 0) 

	order1 := new(Order) 
	priority1, quantity1 := 2, 20 
	product1, customerName1 := "Computer", "Greg White" 

	order1.NewOrder(priority1, quantity1, product1, customerName1) 

	order2 := new(Order) 
	priority2, quantity2 := 1, 10 
	product2, customerName2 := "Monitor", "John Smith" 

	order2.NewOrder(priority2, quantity2, product2, customerName2) 
	queue.Add(order1) 
	queue.Add(order2) 

	for i:= 0; i < len(queue); i++{
		fmt.Println(queue[i])
	}
}