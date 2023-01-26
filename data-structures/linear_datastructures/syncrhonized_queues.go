package main 

import (
	"fmt" 
	"math/rand" 
	"time" 
)

// constants
const (
	messagePassStart = iota 
	messageTicketStart 
	messagePassEnd 
	messageTicketEnd
)

type Queue struct{
	waitPass int  
	waitTicket int 
	playPass bool 
	playTicket bool 
	queuePass chan int 
	queueTicket chan int 
	message chan int 
}

func (q *Queue) NewQeue(){
	q.message = make(chan int) 
	q.queuePass = make(chan int) 
	q.queueTicket = make(chan int) 

	go func(){
		var message int 
		select{
		case message = <-q.message:
			switch message{
			case messagePassStart: 
				q.waitPass++ 
			case messagePassEnd: 
				q.playPass = false 
			case messageTicketStart: 
				q.waitTicket++ 
			case messageTicketEnd: 
				q.playTicket = false 

			}
			if q.waitPass > 0 && q.waitTicket > 0 && !q.playPass && !q.playTicket{
				q.playPass = true 
				q.playTicket = true 
				q.waitTicket-- 
				q.waitPass-- 
				q.queuePass <- 1 
				q.queueTicket <- 1
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue 
func (q *Queue) StartTicketIssue(){
	q.message <- messageTicketStart 
	<-q.queueTicket
}

// EndTicketIssue ends the ticket issue 
func (q *Queue) EndTicketIssue(){
	q.message <- messageTicketEnd 
}

// StartPass ends Pass queue 
func (q *Queue) StartPass(){
	q.message <- messagePassStart 
	<-q.queuePass
}

func (q *Queue) EndPass(){
	q.message <- messagePassEnd 
}

func ticketIssue(q *Queue){
	for{
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond) 
		q.StartTicketIssue() 
		fmt.Println("Ticket Issue Starts") 

		// Sleep up to 2 seconds 
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) 
		fmt.Println("Ticket issue ends") 
		q.EndTicketIssue() 
	}
}

// passenger method starts and ends the pass queue 
func passenger(q *Queue){
	for{
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond) 
		q.StartPass() 
		fmt.Println(" Passenger starts") 
		// Sleep up to 2 seconds 
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) 
		fmt.Println(" Passenger ends") 
		q.EndPass()
	}
}


func main(){
	q := new(Queue) 
	q.NewQeue() 
	fmt.Println(q) 
	for i:= 0; i < 10; i++{
		go passenger(q)
	}

	for j:= 0; j < 5; j++{
		go ticketIssue(q)
	}
	select {}
}