package main

/**
	The net/rpc (remote procedure call) and net/rpc/jsonrpc packages 
	provide an easy way to expose methods so they can be invoked over 
	a network. (Rather than just in the programm running them)
**/

import (
	"fmt" 
	"net"
	"net/rpc"
)

type Server struct{} 

func (this *Server) Negate(i int64, reply *int64) error{

	*reply = -i 
	return nil
}

func server(){
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999") 
	if err != nil{
		fmt.Println(err) 
		return 
	}
	for {
		c, err := ln.Accept() 
		if err != nil{
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client(){
	c, err := rpc.Dial("tcp", "127.0.0.1:9999") 
	if err != nil{
		fmt.Println(err) 
		return 
	}

	var result int64 
	err = c.Call("Server.Negate", int64(999), &result) 
	if err != nil{
		fmt.Println(err) 
	} else{
		fmt.Println("Server.Negate(999) = ", result)
	}
}

func main(){
	go server() 
	go client() 
	_ , err := net.Listen("tcp", ":9999") 
	if err != nil{
		fmt.Println("Server, Listen", err)
	}
	var input string 
	fmt.Scanln(&input)
}